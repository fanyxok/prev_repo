classdef Server < handle
    %斗地主相关数据的存储中心，
    %   此处显示详细说明
    
    properties(SetAccess = private)
        room_status_
        category_
        controller_
        CARDS
        
        accounts_
        points_
    end
    
    methods
        function self = Server()
            %构造函数 
            
            self.accounts_ = containers.Map();
            self.points_ = containers.Map();
           
            self.room_status_ = [];
            self.category_ = [];
            self.controller_ = GameController(self);
            self.CARDS = Cards();
        end
        
        function ret = cards(self)
            ret = self.CARDS;
        end
        
        function ret = controller(self)
            ret = self.controller_;
        end
        function ret = points(self, name)
            is_exist = self.points_.isKey(name);
            if (is_exist)
                ret = self.points_(name);
            else
                ret = 0;
            end
            
        end
       %=====================Account========================== 
        function ret = checkAccountExistance(self,name)
            %检查账户name是不是已经存在
            ret = true;
            if (self.accounts_.isKey(name) == 0)
                ret = false;
                
            end 
        end
        
        function ret = addAccount(self, name, password)
            %创建新账户，并记录入库
            ret = true;
            is_exist = self.checkAccountExistance(name);
            if (~is_exist)
                self.accounts_(name) = password;
                self.points_(name) = 0;
                save('Server.mat', 'self'); 
            else
                msgbox('Username already exist');
                ret = false;
            end
        end
        
        function ret = login(self, name, word)
            ret = false;
            is_exist = self.checkAccountExistance(name);
            if (is_exist)
                if (strcmp(self.accounts_(name), word))
                    ret = true;
                else
                    errordlg('Wrong password');
                end
            else
                errordlg('Username not exist');
            end
        end
        
        function deleteAccount(self,name)
            self.accounts_.remove(name);
            self.points_.remove(name);
        end
        
        function updateplayerPoints(self, player, var)
            value = self.points_(player);
            self.points_(player)= value + var;
        end
        
    end
end

