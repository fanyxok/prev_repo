#ifndef PERCEPTRON_BRANCH_PREDICTION
#define PERCEPTRON_BRANCH_PREDICTION


#include "branch_predictor.h"

#include <vector>

class PerceptronBranchPredictor : public BranchPredictor 
{

public:
    struct Perceptron {
        int w0;
        std::vector<int> weights;
    };
    typedef PerceptronBranchPredictor::Perceptron Perceptron;

public:
    PerceptronBranchPredictor(String name, core_id_t core_id, UInt32 size);
    ~PerceptronBranchPredictor();

    bool predict(IntPtr raw, IntPtr target);
    void update(bool predicted, bool actual, IntPtr ip, IntPtr target );

private:
    std::vector<Perceptron> m_perceptrons;
    std::vector<int> m_history;
    int m_history_length;
    int m_perceptron_table_size;
    int m_threshold;
    int m_last_out;

private:
    int hash(IntPtr key); // hash branch address to entry index
    Perceptron& select(int index);
    void updateHistory(bool actual);
    void initWeights(Perceptron &p);
    void initHistory();
    int outcome( Perceptron &);
    int cut(int );
    void train( Perceptron &p, int t);
};










#endif