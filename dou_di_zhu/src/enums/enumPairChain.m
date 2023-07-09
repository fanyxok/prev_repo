function ret = enumPairChain(len)
%枚举连对， 长度大于等于3
    load('CARDS.mat');
    chain_base = [3,10];
    chain_total_number = chain_base(2) - (len - chain_base(1));
    ret = cell(2, chain_total_number);
    for i=1:chain_total_number
        ret(1,i) = { strjoin(CARDS_PAIR(i:i+len-1),"-") };
        ret(2,i) = {CARDS2RANK(CARDS(i))};
    end
    [~,size_t] = size(ret);
    for i = 1:size_t
        ret{1,i} = strjoin(sortCards(chain2cards(ret{1,i})),"-");
    end
end
