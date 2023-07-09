#include <iostream>
#include <fstream>
#include <string>

#include "simulator.h"
#include "perceptron_branch_predictor.h"

void print(std::string s) {
    std::cout << s  << std::endl;
}

// initial parameter, threshold is rely on the length of history
PerceptronBranchPredictor::PerceptronBranchPredictor(String name, core_id_t core_id, UInt32 size)
    : BranchPredictor(name, core_id),
    m_history_length(62),
    m_perceptron_table_size(256),
    m_threshold(static_cast<int>(1.93 * m_history_length + 14))

{
    // create perceptron table
    m_perceptrons = std::vector<Perceptron>(m_perceptron_table_size);

    // init weights and history as 0.
    for ( int i = 0; i < m_perceptrons.size(); i++ ) {
        m_perceptrons[i].weights = std::vector<int>(m_history_length);
        initWeights((m_perceptrons[i]));
    }
    m_history = std::vector<int>(m_history_length);
    initHistory();

    print("Consturt Function Complete");
}

PerceptronBranchPredictor::~PerceptronBranchPredictor() {}

void PerceptronBranchPredictor::initWeights(Perceptron& p){
    int w = 0;
    int l = 0;
    p.w0 = w;
    for (int i = 0; i < p.weights.size(); i++) {
        p.weights[i] = l;
    }
}

void PerceptronBranchPredictor::initHistory(){
    for (int i = 0; i < m_history.size(); i++) {
        m_history[i] = 0;
    }

    // std::string line;
    // std::ifstream sample ("sample.txt");
    // if (sample.is_open())
    // {
    //     int i = 0;
    //     while ( getline (sample,line) && ( i < m_history_length ) )
    //     {
    //         int k = std::stoi(line);
    //         m_history[i] = k;
    //         i++;
            
    //     }
    //     sample.close();
    // }

}

PerceptronBranchPredictor::Perceptron& PerceptronBranchPredictor::select(int index){
    return m_perceptrons[index];
}

// update history by erase the beginning history and pad it with last history
void PerceptronBranchPredictor::updateHistory(bool actual){
    m_history.erase(m_history.begin());
    m_history.push_back(actual ? 1 : -1);
}

// ignore the last 4 bits, and mod by perceptron table size
int PerceptronBranchPredictor::hash(IntPtr key) {
    return ((key >> 4) % m_perceptron_table_size);
}

bool PerceptronBranchPredictor::predict(IntPtr raw, IntPtr target) 
{
    int index = hash(raw);
    Perceptron &p = select(index);
    // compute the output of perceptron
    m_last_out = outcome(p);
     
    // if perceptron output > 0, return true(taken), otherwise return false(not taken)
    return m_last_out > 0 ? true : false;
}

void printHistory(std::vector<int>& v) {
    std::cout << std::endl;
    for( int i = 0; i < v.size(); i++) {
        char b[5];
        std::snprintf(b, 4," %2d ",v[i]);
        std::cout << b;
    
    }
    std::cout << std::endl;
}
void PerceptronBranchPredictor::update(bool predicted, bool actual, IntPtr ip, IntPtr target)
{
    updateCounters(predicted, actual);

    int index = hash(ip);
    Perceptron &p = select(index);
    int actual_i =  actual ? 1 : -1;

    // if perceptron give a wrong prediction or it has low confidence, train it
    if ( predicted != actual || abs(m_last_out) <= m_threshold ) {
        train(p, actual_i);
    }
    //printHistory(m_perceptrons[index].history);
    updateHistory(actual);
    //printHistory(m_perceptrons[index].history);
}

// output of perceptron is computed by the affine equation
int PerceptronBranchPredictor::outcome(Perceptron &p){
    int y = p.w0;
    for( int i = 0; i < p.weights.size(); i++) {
        y += (p.weights[i] * m_history[i]);
    }
    return y;
}

int PerceptronBranchPredictor::cut(int y) {
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

// train perceptron follow the paper said
void PerceptronBranchPredictor::train( Perceptron &p, int t) {
    p.w0 = p.w0 + t;
    for ( int i = 0; i < p.weights.size(); i++ ) {
        p.weights[i] += ( t * m_history[i]);
    }
}