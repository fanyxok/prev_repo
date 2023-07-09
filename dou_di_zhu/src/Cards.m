classdef Cards
    %UNTITLED 此处显示有关此类的摘要
    %   此处显示详细说明
    
    properties
        jokers = ["BJ", "CJ"]
        non_jokers = ["A","2","3","4","5","6","7","8","9","10","J","Q","K"]
        rank_str = ["3","4","5","6","7","8","9","10","J","Q","K","A","2","BJ","CJ"]
        rank_int = [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15]
        % s=黑桃 h=红心 d = 方块 c = 梅花
        suit = [ "s","h", "d", "c"]
        cards
        str2rankMap %str to int rank "3" to 1 "CJ" to 15
        CARDS_TYPE
    end
    
    properties (Access=public)

    end
    methods
        function self = Cards()
            self.cards = [self.non_jokers+self.suit(1), self.non_jokers+self.suit(2),...
                self.non_jokers+self.suit(3), self.non_jokers+self.suit(4), self.jokers];
            self.str2rankMap = containers.Map(self.rank_str,self.rank_int);
            c = load('CARDS_TYPE.mat');
            self.CARDS_TYPE = c.CARDS_TYPE;
            
        end
        
        %"3" to 1
        function ret = rankStrs2Ints(self, strs)
            ret = [];
            [~, size_t] = size(strs);
            for i=1:size_t
                ret(end+1) = self.str2rankMap(strs(i));
            end
        end
        %"3?" to 1
        function ret = suit2str(self,suit_strs)
            ret = suit_strs;
            len = length(suit_strs);
            for i = 1:len
                if (self.isJoker(suit_strs(i)))
                    ret(i) = suit_strs(i);
                else
                    ch = char(suit_strs(i));
                    ret(i) = string(ch(1:end-1));
                end
            end
        end
        function ret = rankSuitStr2Int(self, suit_strs)
            ret = [];
            len = length(suit_strs);
            for i = 1 : len
                if (self.isJoker(suit_strs(i)))
                    if (strcmp(suit_strs(i), "BJ"))
                        ret(end+1) = 14;
                    else
                        ret(end+1) = 15;
                    end
                else
                    ch = char(suit_strs(i));
                    str = string(ch(1:end-1));
                    ret(end+1) = self.rankStrs2Ints(str);
                end
            end
        end
        
        function bool = isJoker(self, card)
            k = strcmp(self.jokers, card);
            if (sum(k)==0)
                bool = false;
            else
                bool = true;
            end
        end
        
        function  sorted_suit_str_cards = sortCardsByRank(self,suit_str_cards)
            %sort "3?" type cards
            int_cards = [];
            for i = 1 : suit_str_cards.size(2)
                self.rankSuitStr2Int(suit_str_cards(i))
                int_cards(end+1) = self.rankSuitStr2Int(suit_str_cards(i));
            end
            [~,i] = sort(int_cards, 'descend');
            sorted_suit_str_cards = suit_str_cards(i);
        end
        
        function ret = typeAndWeight(self, cards)
            %input cards set ["A","Q","BJ"]
            ret = [{"illegal"};{}];
            [~, size_t] = size(self.CARDS_TYPE);
            cards_rank = self.rankSuitStr2Int(cards);
            disp(cards_rank);
            [~,I] = sort(cards_rank,'descend');
            cards = self.suit2str(cards);
            cards = cards(I);
            cards_chain = cards2chain(cards);
            disp(cards_chain);
            for i=1:size_t
                this_map = self.CARDS_TYPE{i};
                enum = this_map('enum');
                str = string(enum(1,:));
                is_member = ismember(str, cards_chain);
                if (sum(is_member) > 0)
                    ret{1} = this_map('name');
                    weight = cell2mat(enum(2,:));
                    ret{2} = weight(is_member);
                    break;
                end
            end
        end
         
         
    end
end

