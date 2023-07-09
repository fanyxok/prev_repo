ImageDir='images/';
raw_img = imread([ImageDir '233.png']);
img=imresize(raw_img, 1) ;

[m,n,d] = size(img);

hs = 64;
hr = 32;
thr = 1;

disp('Iter ...');
modes=mean_shift(img,hs,hr,thr) ;
segm_img = cluster(modes,size(img),hs,hr);

figure(1) ;
imagesc(img);
axis image ; axis off ;

figure(2) ;
imagesc(label2rgb(segm_img-1,'spring','c','shuffle')) ;
axis image ; axis off ;

