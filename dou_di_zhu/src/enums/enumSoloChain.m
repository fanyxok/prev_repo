function ret = enumSoloChain(len)
%枚举顺子，长度大于等于5
    load('CARDS.mat');
    chain_base = [5,8]; %at least length 5, with combination number 8
    chain_total_number = chain_base(2) - (len - chain_base(1));
    ret = cell(2,chain_total_number);
    for i = 1:chain_total_number
        ret(1, i) = {strjoin(CARDS(i:i+len-1),"-")};
        ret(2, i) = {CARDS2RANK(CARDS(i))};
    end    
    [~,size_t] = size(ret);
    for i = 1:size_t
        ret{1,i} = strjoin(sortCards(chain2cards(ret{1,i})),"-");
    end
end
