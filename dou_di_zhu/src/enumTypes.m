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
        CARDS_TYPE{end+1} = containers.Map(key, {"solo", "单张",enumSolo, 15});
        CARDS_TYPE{end+1} = containers.Map(key, {"solo_chain_5","顺子5连",enumSoloChain(5), 8});
        CARDS_TYPE{end+1} = containers.Map(key, {"solo_chain_6","顺子6连",enumSoloChain(6), 7});
        CARDS_TYPE{end+1} = containers.Map(key, {"solo_chain_7","顺子7连",enumSoloChain(7), 6});
        CARDS_TYPE{end+1} = containers.Map(key, {"solo_chain_8","顺子8连",enumSoloChain(8), 5});
        CARDS_TYPE{end+1} = containers.Map(key, {"solo_chain_9","顺子9连",enumSoloChain(9), 4});
        CARDS_TYPE{end+1} = containers.Map(key, {"solo_chain_10","顺子10连",enumSoloChain(10), 3});
        CARDS_TYPE{end+1} = containers.Map(key, {"solo_chain_11","顺子11连",enumSoloChain(11), 2});
        CARDS_TYPE{end+1} = containers.Map(key, {"solo_chain_6","顺子12连",enumSoloChain(12), 1});
        
        CARDS_TYPE{end+1} = containers.Map(key, {"pair","对子",enumPair, 13});
        CARDS_TYPE{end+1} = containers.Map(key, {"pair_chain_3","对子3连",enumPairChain(3), 10});
        CARDS_TYPE{end+1} = containers.Map(key, {"pair_chain_4","对子4连",enumPairChain(4), 9});
        CARDS_TYPE{end+1} = containers.Map(key, {"pair_chain_5","对子5连",enumPairChain(5), 8});
        CARDS_TYPE{end+1} = containers.Map(key, {"pair_chain_6","对子6连",enumPairChain(6), 7});
        CARDS_TYPE{end+1} = containers.Map(key, {"pair_chain_7","对子7连",enumPairChain(7), 6});
        CARDS_TYPE{end+1} = containers.Map(key, {"pair_chain_8","对子8连",enumPairChain(8), 5});
        CARDS_TYPE{end+1} = containers.Map(key, {"pair_chain_9","对子9连",enumPairChain(9), 4});
        CARDS_TYPE{end+1} = containers.Map(key, {"pair_chain_10","对子10连",enumPairChain(10), 3});
        
        CARDS_TYPE{end+1} = containers.Map(key, {"trio","三带零",enumTrio, 13});
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_chain_2","三带零2连",enumTrioChain(2), 11});
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_chain_3","三带零3连",enumTrioChain(3), 10});
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_chain_4","三带零4连",enumTrioChain(4), 9});
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_chain_5","三带零5连",enumTrioChain(5), 8});
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_chain_6","三带零6连",enumTrioChain(6), 7});
       
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_solo","三带一",enumTrioSolo, 13});
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_solo_chain_2","三带一2连",enumTrioSoloChain(2), 11});
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_solo_chain_3","三带一3连",enumTrioSoloChain(3), 10});
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_solo_chain_4","三带一4连",enumTrioSoloChain(4), 9});
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_solo_chain_5","三带一5连",enumTrioSoloChain(5), 12512});
        
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_pair","三带二",enumTrioPair, 13});
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_pair_chain_2","三带二2连",enumTrioPairChain(2), 11});
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_pair_chain_3","三带二3连",enumTrioPairChain(3), 10});
        CARDS_TYPE{end+1} = containers.Map(key, {"trio_pair_chain_4","三带二4连",enumTrioPairChain(4), 3726});
        
        CARDS_TYPE{end+1} = containers.Map(key, {"four_dual_solo","四带二单",enumBombDualSolo, 1170});
        CARDS_TYPE{end+1} = containers.Map(key, {"four_dual_pair","四带二对",enumBombDualPair, 858});
        
        CARDS_TYPE{end+1} = containers.Map(key, {"bomb","炸弹",enumBomb, 13});
        CARDS_TYPE{end+1} = containers.Map(key, {"rocket","王炸",enumRocket, 1});
        
        
        
        