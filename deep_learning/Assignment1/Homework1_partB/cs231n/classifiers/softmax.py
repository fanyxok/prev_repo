from builtins import range
import numpy as np
from random import shuffle
from past.builtins import xrange

def softmax_loss_naive(W, X, y, reg):
    """
    Softmax loss function, naive implementation (with loops)

    Inputs have dimension D, there are C classes, and we operate on minibatches
    of N examples.

    Inputs:
    - W: A numpy array of shape (D, C) containing weights.
    - X: A numpy array of shape (N, D) containing a minibatch of data.
    - y: A numpy array of shape (N,) containing training labels; y[i] = c means
      that X[i] has label c, where 0 <= c < C.
    - reg: (float) regularization strength

    Returns a tuple of:
    - loss as single float
    - gradient with respect to weights W; an array of same shape as W
    """
    # Initialize the loss and gradient to zero.
    loss = 0.0
    dW = np.zeros_like(W)

    #############################################################################
    # TODO: Compute the softmax loss and its gradient using explicit loops.     #
    # Store the loss in loss and the gradient in dW. If you are not careful     #
    # here, it is easy to run into numeric instability. Don't forget the        #
    # regularization!                                                           #
    #############################################################################
    # *****START OF YOUR CODE (DO NOT DELETE/MODIFY THIS LINE)*****
    num_class = W.shape[1]
    num_train = X.shape[0]
    for i in range(num_train):
        raw_scores = X[i].dot(W) 
        max_score = np.max(raw_scores)
        scores = raw_scores-max_score # avoid some one equal to 1.
        exp_scores = np.exp(scores)
        sum_exp_scores = np.sum(exp_scores)
        softmax = exp_scores/sum_exp_scores

        loss_i = - (scores[y[i]] - np.log(sum_exp_scores)) 
        loss += loss_i
        for j in range(num_class):

            softmax_j = softmax[j]
            if j == y[i]:
                dW[:,j] += (-1 + softmax_j) *X[i] 
            else: 
                dW[:,j] += softmax_j *X[i] 

    loss /= num_train 
    loss += reg * np.sum(W * W)
    dW = dW/num_train + 2 * reg* W 
    pass

    # *****END OF YOUR CODE (DO NOT DELETE/MODIFY THIS LINE)*****

    return loss, dW


def softmax_loss_vectorized(W, X, y, reg):
    """
    Softmax loss function, vectorized version.

    Inputs and outputs are the same as softmax_loss_naive.
    """
    # Initialize the loss and gradient to zero.
    loss = 0.0
    dW = np.zeros_like(W)

    #############################################################################
    # TODO: Compute the softmax loss and its gradient using no explicit loops.  #
    # Store the loss in loss and the gradient in dW. If you are not careful     #
    # here, it is easy to run into numeric instability. Don't forget the        #
    # regularization!                                                           #
    #############################################################################
    # *****START OF YOUR CODE (DO NOT DELETE/MODIFY THIS LINE)*****
    num_class = W.shape[1]
    num_train = X.shape[0]
    raw_scores = X.dot(W)
    scores = raw_scores - np.max(raw_scores, axis = 1).reshape(-1,1)
    softmax = np.exp(scores)/np.sum(np.exp(scores), axis = 1).reshape(-1,1)
    loss = -np.sum(np.log(softmax[range(num_train), list(y)]))
    loss /= num_train 
    loss +=  reg * np.sum(W * W)

    # for correct, dw = (p-1)x, else dw = (p)x

    dS = softmax 
    dS[range(num_train), list(y)] += -1 
    dW = (X.T).dot(dS)
    dW = dW/num_train + 2 * reg * W 
    pass

    # *****END OF YOUR CODE (DO NOT DELETE/MODIFY THIS LINE)*****

    return loss, dW
