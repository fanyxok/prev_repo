function ret = cards2chain(cards)
    %cards ["A","2","Q"] to "A-2-Q"
    ret = strjoin(cards, "-");
end
