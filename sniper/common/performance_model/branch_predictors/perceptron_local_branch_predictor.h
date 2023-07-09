#ifndef PERCEPTRON_LOCAL_BRANCH_PREDICTION
#define PERCEPTRON_LOCAL_BRANCH_PREDICTION


#include "branch_predictor.h"

#include <vector>

class PerceptronLocalBranchPredictor : public BranchPredictor 
{

public:
    struct PerceptronWithLocalHistory {
        int w0;
        std::vector<int> weights;
        std::vector<int> history;
    };
    typedef PerceptronLocalBranchPredictor::PerceptronWithLocalHistory PerceptronL;

public:
    PerceptronLocalBranchPredictor(String name, core_id_t core_id, UInt32 size);
    ~PerceptronLocalBranchPredictor();

    bool predict(IntPtr raw, IntPtr target);
    void update(bool predicted, bool actual, IntPtr ip, IntPtr target );

private:
    std::vector<PerceptronL> m_perceptrons;
    int m_history_length;
    int m_perceptron_table_size;
    int m_threshold;
    int m_last_out;

private:
    int hash(IntPtr key); // hash branch address to entry index
    PerceptronL& select(int index);
    void updateHistory(bool actual, IntPtr ip);
    void initWeights(PerceptronL &p);
    void initHistory(PerceptronL &p);
    int outcome( PerceptronL &);
    int cut(int );
    void train( PerceptronL &p, int t);
};










#endif