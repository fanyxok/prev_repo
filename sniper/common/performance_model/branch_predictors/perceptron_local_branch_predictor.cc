#include "simulator.h"
#include "perceptron_local_branch_predictor.h"

// void print(std::string s) {
//     std::cout << s  << std::endl;
// }

PerceptronLocalBranchPredictor::PerceptronLocalBranchPredictor(String name, core_id_t core_id, UInt32 size)
    : BranchPredictor(name, core_id),
    m_history_length(62),
    m_perceptron_table_size(256),
    m_threshold(static_cast<int>(1.93 * m_history_length + 14))

{
    m_perceptrons = std::vector<PerceptronL>(m_perceptron_table_size);

    for ( int i = 0; i < m_perceptrons.size(); i++ ) {
        m_perceptrons[i].weights = std::vector<int>(m_history_length);
        m_perceptrons[i].history = std::vector<int>(m_history_length);
        initWeights((m_perceptrons[i]));
        initHistory((m_perceptrons[i]));
    }
    //print("Consturt Function Complete");
}

PerceptronLocalBranchPredictor::~PerceptronLocalBranchPredictor() {}

void PerceptronLocalBranchPredictor::initWeights(PerceptronL& p){
    int w = 0;
    int l = 0;
    p.w0 = w;
    for (int i = 0; i < p.weights.size(); i++) {
        p.weights[i] = l;
    }
}

void PerceptronLocalBranchPredictor::initHistory(PerceptronL& p){
    for (int i = 0; i < p.history.size(); i++) {
        p.history[i] = 0;
    }
}

PerceptronLocalBranchPredictor::PerceptronL& PerceptronLocalBranchPredictor::select(int index){
    return m_perceptrons[index];
}

void PerceptronLocalBranchPredictor::updateHistory(bool actual, IntPtr ip){
    int index = hash(ip);
    PerceptronL &p = select(index);
    p.history.erase(p.history.begin());
    p.history.push_back(actual ? 1 : -1);
}

int PerceptronLocalBranchPredictor::hash(IntPtr key) {

    return ((key >> 4) % m_perceptron_table_size);
}

bool PerceptronLocalBranchPredictor::predict(IntPtr raw, IntPtr target) 
{
    int index = hash(raw);
    PerceptronL &p = select(index);
    m_last_out = outcome(p);
     
    return m_last_out > 0 ? true : false;
}

// void printHistory(std::vector<int>& v) {
//     std::cout << std::endl;
//     for( int i = 0; i < v.size(); i++) {
//         char b[5];
//         std::snprintf(b, 4," %2d ",v[i]);
//         std::cout << b;
    
//     }
//     std::cout << std::endl;
// }
void PerceptronLocalBranchPredictor::update(bool predicted, bool actual, IntPtr ip, IntPtr target)
{
    updateCounters(predicted, actual);

    int index = hash(ip);
    PerceptronL &p = select(index);
    int actual_i =  actual ? 1 : -1;

    if ( predicted != actual || abs(m_last_out) <= m_threshold ) {
        train(p, actual_i);
    }
    //printHistory(m_perceptrons[index].history);
    updateHistory(actual, ip);
    //printHistory(m_perceptrons[index].history);
}

int PerceptronLocalBranchPredictor::outcome(PerceptronL &p){
    int y = p.w0;
    for( int i = 0; i < p.weights.size(); i++) {
        y += (p.weights[i] * p.history[i]);
    }
    return y;
}

int PerceptronLocalBranchPredictor::cut(int y) {
    int out;
    if ( y > m_threshold ) {
        out = 1;
    }else if ( y < m_threshold ) {
        out = -1;
    }else {
        out = 0;
    }
    return out;
}

void PerceptronLocalBranchPredictor::train( PerceptronL &p, int t) {
    p.w0 = p.w0 + t;
    for ( int i = 0; i < p.weights.size(); i++ ) {
        p.weights[i] += ( t * p.history[i]);
    }
}