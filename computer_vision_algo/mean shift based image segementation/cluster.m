function segm = cluster(modes_map,size_i,hs,hr)
m = size_i(1);
n = size_i(2);

chess = ones( 2*m+1, 2*n+1);
chess(1:2:(2*m+1),:) = zeros( m+1, 2*n+1);
chess(:,1:2:(2*n+1)) = zeros( 2*m+1, n+1);
%cluster row
grad = abs(modes_map(:,2:end,1:2)-modes_map(:,1:(end-1),1:2));
%grad = grad(:,:,1).*grad(:,:,1)+grad(:,:,2).*grad(:,:,2);
%grad = grad.^(0.5)./2;
space_cluster =  grad < (hs*0.6);

grad = abs(modes_map(:,2:end,3:end)-modes_map(:,1:(end-1),3:end));
%grad = grad(:,:,1).*grad(:,:,1)+grad(:,:,2).*grad(:,:,2)+grad(:,:,3).*grad(:,:,3);
%grad = grad.^(0.5)./3;
color_cluster = grad < (hr);
is_nearby = cat(3,space_cluster,color_cluster);
chess(2:2:2*m,3:2:(2*n-1)) = all(is_nearby, 3); 

%cluster col
space_cluster = abs(modes_map(2:end,:,1:2)-modes_map(1:(end-1),:,1:2)) < (hs);
color_cluster = abs(modes_map(2:end,:,3:end)-modes_map(1:(end-1),:,3:end)) < (hr);
is_nearby = cat(3, space_cluster, color_cluster);
chess(3:2:(2*m-1),2:2:2*n) = all(is_nearby, 3);

segm = bwlabel( chess, 4 );    
segm = segm( 2:2:2*m, 2:2:2*n );    
end

