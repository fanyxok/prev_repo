%枚举所有飞机不带翼(连三不带)，最少2连
function ret = enumTrioChain(len)
    load('CARDS.mat');
    chain_base = [2, 11];
    total_number = chain_base(2)-(len-chain_base(1));
    ret = cell(2, total_number);
    for i = 1 : total_number
        ret(:, i) = [ { strjoin(CARDS_TRIO(i:i+len-1), "-") } , { CARDS_RANK(i) } ];
    end
end
