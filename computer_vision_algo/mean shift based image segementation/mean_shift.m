function result_mode=mean_shift(img,hs,hr,thr) 

[m,n,d] = size(img);
img = double(img);
% h is [hs, hs, hr, hr, hr]
h  = [hs hs repmat(hr,1,d)];
% a_conv_b contains conv info that pixel a conv to pixel b and the info for
% color space
a_conv_b  = zeros( m, n, d+2 );
for x = 1:n
  for y = 1:m
     
    %reshape y to [position x, position y, d1, d2, d3 ]
    pixel = double( [x y reshape(img(y,x,:),1,d)] );

    x_left = max(x-hs,1);  
    x_right = min(x+hs,n);  
    length = x_right-x_left+1;
    
    y_up = max(y-hs,1);  
    y_down = min(y+hs,m);  
    hight = y_down-y_up+1;
    
    squares = length*hight;  
    %iw is a 1 X nw matrix, value from 0 to nw-1.
    serial = (0:(squares-1))';
    position_xs = floor(serial/hight)+x_left;
    position_ys = mod(serial,hight)+y_up;
    %change interset region [height * length * 3] to [(height*length) * 3]
    reshape_= reshape(img(y_up:y_down,x_left:x_right,:),[],d);
    %region_sr is (height*length)*5 which contains [position of each pixel of the region, d1, d2, d3]
    region_sr = [position_xs position_ys reshape_];
    
    %compute mean shift, current pixel conv to whick pixel 
    while true
        %compute (|x_i[s,r]-x[s,r]|/h[s,r])**2 lable as M 
        g_para = (region_sr-repmat(pixel,squares,1)) ./ repmat(h,squares,1);
        g_para = sum(g_para.*g_para, 2);
        
        %compute g(M), g(x) is 1 if x leq 1,otherwise 0. 
        %g = g_para<1.0;
        record = pixel;
        
        K = (2*pi)^(-0.5)*exp(-g_para./2);
        sumK=sum(K);
        BigK = repmat(K,1,(d+2));
        pixel = sum(region_sr(:,:).*BigK)/sumK;
        
        %compute value of mean shift = sum(x_i[s,r].*g(M))/sum(g(M))
        %pixel = mean( region_sr(g,:), 1 );
        step = norm(pixel-record);
        if step < thr * max(hs,hr)
            break; 
        end
    end
    a_conv_b(y,x,:) = pixel;
  end 
end 
result_mode = a_conv_b;              
