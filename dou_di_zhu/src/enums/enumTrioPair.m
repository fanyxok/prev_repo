
function ret = enumTrioPair()
%枚举三带二
load('CARDS.mat');
    size_t = CARDS_TRIO.size(2);
    size_rank = CARDS.size(2);
    total_number = size_t*(size_rank-2);
    ret = cell(2, total_number);
    ed = 1;
    for i = 1:size_t
        for j = 1:(size_rank-1)
            %this = trio(i)+"-"+CARDS(j);
            if (i == j) 
                continue
            end
            if (j == (size_rank-1))
                ret(:,ed) = [{CARDS_TRIO(i)+"-"+CARDS(j)+"-"+CARDS(j+1)}, {CARDS_RANK(i)}];
                ed=ed+1;
            else
                ret(:,ed) = [{CARDS_TRIO(i)+"-"+CARDS(j)+"-"+CARDS(j)}, {CARDS_RANK(i)}];
                ed=ed+1;
            end
        end
    end
    [~,size_t] = size(ret);
    for i = 1:size_t
        ret{1,i} = strjoin(sortCards(chain2cards(ret{1,i})),"-");
    end
end