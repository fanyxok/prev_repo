function ret = enumSolo()
    %ö���������ͣ�������Ȩ�أ��Ա�Ƚϴ�С
    %ret = (weight, cards)
    load('CARDS.mat');
    [~,size_t] = size(CARDS);
    ret = cell(2, size_t);
    card_cell = cellstr(CARDS);
    weight_cell = num2cell(CARDS_RANK);
    ret(1,:) = card_cell;
    ret(2,:) = weight_cell;
end
