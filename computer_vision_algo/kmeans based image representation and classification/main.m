run('vlfeat-0.9.21-bin/vlfeat-0.9.21/toolbox/vl_setup');

%--------------SIFT train data---------------
%提取data set中所有图片的SIFT特征,并保存每张图的feature数量信息
imgdir = textread('imgdirs.txt', '%s');
[class_number,~] = size(imgdir);
train_number = 30;
[sift_d_mat, feature_num] = do_sifts(class_number, train_number, imgdir);

%------------Kmeans----------------------
K = 500;
%do K-Means
[Idx, C] = kmeans(single(sift_d_mat'),K);
%----------counter and save histgram
counter_l = 1;
counter_r = 0;
imgs2cluster = zeros(class_number*train_number, K);
for i=1:class_number
    for j=1:train_number
        number = (i-1)*train_number+j;
        counter_r = counter_l + feature_num(number);
        [countM, ~] = hist(Idx(counter_l:counter_r-1), 1:K); 
        counter_l = counter_r;
        imgs2cluster(number, :) = countM;
    end
end
clear countM counter_l counter_r number
%-------------train--------------------        
train_lable_vector = repmat([1:class_number],train_number,1);
train_lable_vector = reshape(train_lable_vector,[],1);
train = svmtrain(train_lable_vector, imgs2cluster);


%------------load and sift test images-----------
test_number = 10;
[sift_d_mat_t, feature_num_t] = do_sifts2(class_number, train_number,test_number, imgdir);

%-----------features2cluster vector------------
[m_t, n_t] = size(sift_d_mat_t);
%compute L2 distance
dist_min = inf;
dist_min_idx = 0;
imgs_features2cluster = [];
for i=1:n_t
    curr_sift = sift_d_mat_t(:,i);
    rept_curr_sift = repmat(curr_sift,1, K);
    dists = single(rept_curr_sift)-C';
    dist = sqrt(sum(dists.*dists,1));
    dist_min = (find(dist==min(dist)));
    imgs_features2cluster = [imgs_features2cluster dist_min(1)];
end
%--------------imgs2clusters-------------------
counter_l = 1;
counter_r = 0;
imgs2cluster_t = [];
for i=1:class_number*test_number
    counter_r = counter_l + feature_num_t(i);
    [hist_m, ~] = hist(imgs_features2cluster(counter_l:counter_r-1), 1:K);
    counter_l = counter_r;
    imgs2cluster_t(i,:) = hist_m;
end

%-----------tfidf----------------------
[size_m, size_n] = size(imgs2cluster_t);
%df, documents are imgs
DF_t = [];
for i=1:K
    d = sum(imgs2cluster_t(:,i) > 0);
    DF_t = [DF_t d];
end
%tfidf 
tfidf_t = [];
for i=1:size_m
    min_f = min(imgs2cluster_t(i,:));
    max_f = max(imgs2cluster_t(i,:));
    for j=1:K
        tf = (imgs2cluster_t(i,j)-min_f)/(max_f-min_f);
        tfidf_t(i,j) = log(tf+1) * log(size_m/(DF_t(j)+1));
    end
end
%--------------test-------------------
test_lable_vector = repmat([1:class_number],test_number,1);
test_lable_vector = reshape(test_lable_vector,[],1);
predict = svmpredict(test_lable_vector, imgs2cluster_t, train);
