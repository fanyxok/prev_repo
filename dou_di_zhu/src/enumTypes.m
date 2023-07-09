        %pre-define cards set
        CARDS = ["3","4","5","6","7","8","9","10","J","Q","K","A","2","BJ","CJ"];
        CARDS_RANK = [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15];
        CARDS2RANK = containers.Map(CARDS,CARDS_RANK);
        CARDS_NO_JOKERS = ["3","4","5","6","7","8","9","10","J","Q","K","A","2"];
        CARDS_JOKERS =  ["BJ", "CJ"];
        CARDS_PAIR = CARDS_NO_JOKERS + "-" + CARDS_NO_JOKERS;
        CARDS_TRIO = CARDS_PAIR + "-" + CARDS_NO_JOKERS;
        CARDS_FOUR = CARDS_TRIO + "-" + CARDS_NO_JOKERS;
        
        
        key = ["name", "zh_name", "enum", "size"];
        CARDS_TYPE= {};
        CARDS_TYPE{end+1} = containers.Map(key, {"solo", "����",enumSolo, 15});
        CARDS_TYPE{end+1} = containers.Map(key, {"solo_chain_5","˳��5��",enumSoloChain(5), 8});
        CARDS_TYPE{end+1} = containers.Map(key, {"solo_chain_6","˳��6��",enumSoloChain(6), 7});
        CARDS_TYPE{end+1} = containers.Map(key, {"solo_chain_7","˳��7��",enumSoloChain(7), 6});
        CARDS_TYPE{end+1} = containers.Map(key, {"solo_chain_8","˳��8��",enumSoloChain(8), 5});
        CARDS_TYPE{end+1} = containers.Map(key, {"solo_chain_9","˳��9��",enumSoloChain(9), 4});
        CARDS_TYPE{end+1} = containers.Map(key, {"solo_chain_10","˳��10��",enumSoloChain(10), 3});
        CARDS_TYPE{end+1} = containers.Map(key, {"solo_chain_11","˳��11��",enumSoloChain(11), 2});
        CARDS_TYPE{end+1} = containers.Map(key, {"solo_chain_6","˳��12��",enumSoloChain(12), 1});
        
        CARDS_TYPE{end+1} = containers.Map(key, {"pair","����",enumPair, 13});
        CARDS_TYPE{end+1} = containers.Map(key, {"pair_chain_3","����3��",enumPairChain(3), 10});
        CARDS_TYPE{end+1} = containers.Map(key, {"pair_chain_4","����4��",enumPairChain(4), 9});
        CARDS_TYPE{end+1} = containers.Map(key, {"pair_chain_5","����5��",enumPairChain(5), 8});
        CARDS_TYPE{end+1} = containers.Map(key, {"pair_chain_6","����6��",enumPairChain(6), 7});
        CARDS_TYPE{end+1} = containers.Map(key, {"pair_chain_7","����7��",enumPairChain(7), 6});
        CARDS_TYPE{end+1} = containers.Map(key, {"pair_chain_8","����8��",enumPairChain(8), 5});
        CARDS_TYPE{end+1} = containers.Map(key, {"pair_chain_9","����9��",enumPairChain(9), 4});
        CARDS_TYPE{end+1} = containers.Map(key, {"pair_chain_10","����10��",enumPairChain(10), 3});
        
        CARDS_TYPE{end+1} = containers.Map(key, {"trio","������",enumTrio, 13});
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_chain_2","������2��",enumTrioChain(2), 11});
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_chain_3","������3��",enumTrioChain(3), 10});
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_chain_4","������4��",enumTrioChain(4), 9});
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_chain_5","������5��",enumTrioChain(5), 8});
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_chain_6","������6��",enumTrioChain(6), 7});
       
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_solo","����һ",enumTrioSolo, 13});
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_solo_chain_2","����һ2��",enumTrioSoloChain(2), 11});
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_solo_chain_3","����һ3��",enumTrioSoloChain(3), 10});
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_solo_chain_4","����һ4��",enumTrioSoloChain(4), 9});
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_solo_chain_5","����һ5��",enumTrioSoloChain(5), 12512});
        
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_pair","������",enumTrioPair, 13});
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_pair_chain_2","������2��",enumTrioPairChain(2), 11});
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_pair_chain_3","������3��",enumTrioPairChain(3), 10});
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_pair_chain_4","������4��",enumTrioPairChain(4), 3726});
        
        CARDS_TYPE{end+1} = containers.Map(key, {"four_dual_solo","�Ĵ�����",enumBombDualSolo, 1170});
        CARDS_TYPE{end+1} = containers.Map(key, {"four_dual_pair","�Ĵ�����",enumBombDualPair, 858});
        
        CARDS_TYPE{end+1} = containers.Map(key, {"bomb","ը��",enumBomb, 13});
        CARDS_TYPE{end+1} = containers.Map(key, {"rocket","��ը",enumRocket, 1});
        
        
        
        