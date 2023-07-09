function ret = enumSolo()
    %枚举所有牌型，并附加权重，以便比较大小
    %ret = (weight, cards)
    load('CARDS.mat');
    [~,size_t] = size(CARDS);
    ret = cell(2, size_t);
    card_cell = cellstr(CARDS);
    weight_cell = num2cell(CARDS_RANK);
    ret(1,:) = card_cell;
    ret(2,:) = weight_cell;
end
