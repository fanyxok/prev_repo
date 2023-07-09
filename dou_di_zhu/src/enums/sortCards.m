function ret = sortCards(cards)
%SORTCARDS 此处显示有关此函数的摘要
%   此处显示详细说明
load('CARDS.mat')
    ranks = [];
    [~, size_t] = size(cards);
    for i=1:size_t
        ranks(end+1) = CARDS2RANK(cards(i));
    end
    ranks = sort(ranks,'descend');
    ret = CARDS(ranks);
end

