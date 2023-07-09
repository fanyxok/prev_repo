classdef GameController < handle
    %UNTITLED2 此处显示有关此类的摘要
    %   此处显示详细说明
    
    properties
        server_address 
        room_id %valid id from 1 to 10
        packed = 0 % 3 players? yes 1 no 0     
        room_status % waiting(0) or playing(1)
        player_cards
        server
        CARDS = ["3","4","5","6","7","8","9","10","J","Q","K","A","2","BJ","CJ"];
        CARDS_RANK = [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15];
        CARDS2RANK
        SUIT= ["3s","3h","3d","3c","4s","4h","4d","4c","5s","5h","5d","5c",...
                "6s","6h","6d","6c","7s","7h","7d","7c","8s","8h","8d","8c",...
                "9s","9h","9d","9c","10s","10h","10d","10c","Js","Jh","Jd","Jc",...
                "Qs","Qh","Qd","Qc","Ks","Kh","Kd","Kc","As","Ah","Ad","Ac",...
                "2s","2h","2d","2c","BJ","CJ"];
    end
    
    properties(Access = public)
        last_enter
        last_exit
        players % at most 3
        player_status % 3 status of 3 players, ready(0) or playing(1) or non-ready(-1)  
        dizhu_cards
        player_identity
        bid_points
        bid_times
        SUITED2RANK
        
        largest_who  = []
        last_typeAndWeight = [{"il"};{}];
        curr_player
        last_play = cell(1,3);
     
    end
    
    events
        PlayerEnter
        PlayerExit
        ReadyChanged
        StartGame  
        LandlordAppear
        Play
        Next
        Win
    end

    methods
        function self = GameController(server)
            self.room_id = 0;
            self.players = ["No Player","No Player","No Player"];
            self.player_status = ["No Ready","No Ready","No Ready"];
            self.room_status = 0;
            self.packed = 0;
            self.player_cards = cell(1,3);
            self.bid_times = 0;
            self.bid_points = [0,0,0];
            self.player_identity = ["Farmer", "Farmer", "Farmer"];
            self.server = server;
            
            self.CARDS2RANK= containers.Map(self.CARDS,self.CARDS_RANK);
            rank = [reshape([self.CARDS_RANK(1:end-2);self.CARDS_RANK(1:end-2);self.CARDS_RANK(1:end-2);self.CARDS_RANK(1:end-2)], 1 , []),...
                14,15];
            self.SUITED2RANK= containers.Map(self.SUIT,rank);
        end
        
        function startGame(self)
            self.room_status = 1;
            %suit = self.SUIT(randperm(length(self.SUIT)));
            suit = self.SUIT;
            self.player_cards{1} = self.sortCards(suit(1:17));
            self.player_cards{2} = self.sortCards(suit(18:34));
            self.player_cards{3} = self.sortCards(suit(35:51));  
            self.dizhu_cards = suit(52:54);
            self.notify('StartGame');
        end
        
        function endGame(self)
            self.room_status = 0;
            self.bid_times = 0;
            self.bid_points = [0,0,0];
            self.player_identity = ["Farmer", "Farmer", "Farmer"];
            

        end
        
        function updatePoints(self)
            farmer = find(self.player_identity=="Farmer");
            lord= find(self.player_identity=="Landlord");
            k = (self.players==self.curr_player);
            disp(farmer)
            disp(lord)
            disp(k)
            if (strcmp(self.player_identity(k), "Farmer"))
                 self.server.updateplayerPoints(self.players(farmer(1)), 100);
                 self.server.updateplayerPoints(self.players(farmer(2)), 100);
                 self.server.updateplayerPoints(self.players(lord),-100);
            else
                self.server.updateplayerPoints(self.players(farmer(1)), -100);
                self.server.updateplayerPoints(self.players(farmer(2)), -100);
                self.server.updateplayerPoints(self.players(lord),100);
            end
        end
            
            
        %==========================================
        function ret = sortCards(self, suited_cards)
            ranks = arrayfun(@(x) self.SUITED2RANK(x), suited_cards);
            [~,I] = sort(ranks,'descend');
            ret = suited_cards(I);
        end
        
        function ret = getCards(self,name)
            k = strcmp(self.players, name);
            ret = self.player_cards{k};
        end
        
        function removeCards(self,name,idxs)
            k = strcmp(self.players, name);
            self.player_cards{k}(idxs) = [];
        end
        
        function ret = cardsLeft(self,name)
            k = strcmp(self.players, name);
            [~,ret] = size(self.player_cards{k});
        end
        function ret = lastCard(self)
            ret = self.last_typeAndWeight;
        end
        
        function ret = lastCardWho(self)
            ret = self.largest_who;
        end
        
        function setLastPlay(self,name,cards)
             k = strcmp(self.players, name);
             self.last_play{k} = cards;
        end
        
        function ret = playIdx(self,name) 
            ret = strcmp(self.players, name);
        end
        %=============Play Flow Control======================
        function passPlay(self)
            self.setLastPlay(self.curr_player,[]);
            self.notify('Play');
            self.goToNext();
        end
            
            
        function ret = tryPlay(self,idxs)
            ret = false;
            disp(idxs);
            if (~isempty(idxs) == 0)
                ret = false;     
                return;
            end
            
            cards = self.getCards(self.curr_player);
            cards = cards(idxs);
            type_weight = self.server.cards().typeAndWeight(cards);
            type = type_weight{1};
            if ( strcmp(type, 'illegal'))  
                return;
            end
            disp(type_weight);
            weight = type_weight{2};
            if (strcmp(self.curr_player, self.lastCardWho()))         
                if (size(type)>0)
                    ret = true;                  
                end
            elseif strcmp(type, self.last_typeAndWeight{1})
                if (weight > self.last_typeAndWeight{2})
                    ret = true;
                end
            elseif strcmp(type, 'bomb')
                ret = true;
                if (strcmp(self.last_typeAndWeight{1}, 'rocket'))
                    ret = false;
                end
            elseif strcmp(type, 'rocket')
                ret = true;
            end  
            disp(ret)
            if ( ret == true )
                self.last_typeAndWeight = type_weight;
                self.largest_who = self.curr_player;
                %remove play cards from 手牌
                self.removeCards(self.curr_player, idxs);
                %TODO  show play card on play bord
                self.setLastPlay(self.curr_player,cards);
                self.notify('Play');
                %call next
                self.goToNext();
            end
        end
        
        function goToNext(self)
            %update curr_player to next one
            this = self.playIdx(self.curr_player);
            k = find(this==1);
            if (isempty(self.player_cards{k}))
                % notify game end 
               
               self.updatePoints();
               self.notify('Win');
               self.endGame();
               return;
            end
            if (k < 3)
                self.curr_player=self.players(k+1);
            else
                self.curr_player=self.players(1);
            end
            % notify view
            self.notify('Next');
            
        end
        %======================================
        function bid(self, name, point)
            k = strcmp(self.players, name);
            self.bid_points(k) = point;
            %3个人都bid了，开始决定谁是地主
            self.bid_times = self.bid_times+1;
            if (self.bid_times == 3)
                self.setlandlord();
            end
        end
              
        %决定谁是地主，并且把剩下3张牌给他
        function setlandlord(self)
            if (self.bid_points(1) == 1)
                self.player_identity(1) = "Landlord";
                self.player_cards{1} = self.sortCards([self.player_cards{1}, self.dizhu_cards]);
                self.curr_player = self.players(1);
                self.largest_who = self.players(1);
            elseif (self.bid_points(2) == 1)
                self.player_identity(2) = "Landlord";
                self.player_cards{2} =  self.sortCards([self.player_cards{2}, self.dizhu_cards]);
                self.curr_player = self.players(2);
                self.largest_who = self.players(2);
            else
                self.player_identity(3) = "Landlord";
                self.player_cards{3} =  self.sortCards([self.player_cards{3}, self.dizhu_cards]);
                self.curr_player = self.players(3);
                self.largest_who = self.players(3);
            end
            self.notify('LandlordAppear');
        end
        
        function ret = number(self)
            k = (strcmp(self.players, "No Player"));
            ret = 3-sum(k);
        end
        %===========================================
        function ready(self, name)
            k = (strcmp(self.players, name));
            self.player_status(k) = "Ready";
            self.notify('ReadyChanged');
            if (sum(strcmp(self.player_status, "Ready")) == 3)
                self.startGame(); 
            end
        end
        
        function unready(self,name)
            k = (strcmp(self.players, name));
            self.player_status(k) = "No Ready";
            self.notify('ReadyChanged');
        end
        
        function ret = enterPlayer(self, name)
            ret = false;
            if (self.packed)
                return
            end
            ret = true;
            for i=1:3
                if (self.players(i)=="No Player")
                    self.players(i) = name;
                    self.player_status(i) = "No Ready";
                    self.last_enter=name;
                    self.notify('PlayerEnter');
                    break
                end
            end
            if (sum(ismember(self.players,'No Player')) == 0)
                self.packed = 1;
            end          
        end
        
        function exitPlayer(self,name)
            k = find(strcmp(self.players, name));          
            self.players(k) = "No Player";
            self.player_status(k) = "No Ready";
            self.packed = 0;
            self.last_exit=name;
            self.notify('PlayerExit');
        end 
        
        function cleanPlayers(self)
            self.players = ["No Player","No Player","No Player"];
            self.player_status = ["No Ready","No Ready","No Ready"];
            self.packed = 0;
            self.player_cards = cell(1,3);
            self.bid_times = 0;
            self.bid_points = [0,0,0];
            self.player_identity = ["Farmer", "Farmer", "Farmer"];
        end
    end
end

