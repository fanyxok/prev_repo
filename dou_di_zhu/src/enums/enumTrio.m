
function ret = enumTrio()
%ö��������
    load('CARDS.mat');
    ret = cell(2, CARDS_TRIO.size(2));
    ret = [cellstr(CARDS_TRIO);num2cell(CARDS_RANK(1:13))];
end
