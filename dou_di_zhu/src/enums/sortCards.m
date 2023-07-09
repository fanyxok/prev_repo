function ret = sortCards(cards)
%SORTCARDS �˴���ʾ�йش˺�����ժҪ
%   �˴���ʾ��ϸ˵��
load('CARDS.mat')
    ranks = [];
    [~, size_t] = size(cards);
    for i=1:size_t
        ranks(end+1) = CARDS2RANK(cards(i));
    end
    ranks = sort(ranks,'descend');
    ret = CARDS(ranks);
end

