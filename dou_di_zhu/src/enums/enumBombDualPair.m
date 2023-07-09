%ö���Ĵ����ԣ����Բ�����ͬ������һ�Բ�����BJ+CJ��
function ret = enumBombDualPair()
load('CARDS.mat');
    [~, size_t] = size(CARDS_FOUR);
    comb_number = 66;%������91�����
    total_number = size_t * comb_number;
    ret = cell(2, total_number);
    ed = 1;
    for i = 1:size_t
        residual = CARDS_RANK(1:13);
        residual(i) = [];
        kickers = combnk(residual, 2);
        kickers = unique(sort(kickers,2), 'row');
        k = ismember(kickers, [14, 15], 'row');
        kickers(k,:) = [];
        [extend, ~] = size(kickers);
        ret(1,ed:ed+extend-1) = repmat(cellstr(CARDS_FOUR(i)),1,extend);
        ret(2,ed:ed+extend-1) = {CARDS_RANK(i)};
        for j=1:extend
            ret{1,ed+j-1} = ret{1,ed+j-1} + "-" + strjoin(CARDS_PAIR(kickers(j,:)),"-");
        end
        ed = ed + extend;
    end
    [~,size_t] = size(ret);
    for i = 1:size_t
        ret{1,i} = strjoin(sortCards(chain2cards(ret{1,i})),"-");
    end
end