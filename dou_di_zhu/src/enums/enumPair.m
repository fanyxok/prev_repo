
function ret = enumPair()
%ц╤╬ы╤твс
    load('CARDS.mat');
    ret = cell(2, CARDS_PAIR.size(2));
    ret(1,:) = cellstr(CARDS_PAIR);
    ret(2,:) = num2cell(CARDS_RANK(1:13));
end

