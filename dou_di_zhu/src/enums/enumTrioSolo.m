

function ret = enumTrioSolo()
%枚举三带一
    load('CARDS.mat');
    size_trio = CARDS_TRIO.size(2);
    size_rank = CARDS.size(2);
    total_number = size_trio*(size_rank-1);
    ret = cell(2, total_number);
    ed = 1;
    for i = 1:size_trio
        for j = 1:size_rank
            %this = trio(i)+"-"+CARDS(j);
            if (i == j) 
                continue
            end
            cards = [chain2cards(CARDS_TRIO(i)),CARDS(j)];
            cards = sortCards(cards);
            ret(:,ed) = [{strjoin(cards,"-")}, {CARDS_RANK(i)}];
            ed=ed+1;
        end
    end
end