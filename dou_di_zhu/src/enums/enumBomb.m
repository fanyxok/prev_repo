function ret = enumBomb()
load('CARDS.mat');
    [~, size_t] = size(CARDS_FOUR);
    ret = cell(2, size_t);
    ret(1,:) = cellstr(CARDS_FOUR);
    ret(2, :) = num2cell(CARDS_RANK(1:13));
end