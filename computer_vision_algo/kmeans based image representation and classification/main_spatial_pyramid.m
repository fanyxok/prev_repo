run('vlfeat-0.9.21-bin/vlfeat-0.9.21/toolbox/vl_setup');

%--------------SIFT train data---------------
%提取data set中所有图片的SIFT特征,并保存每张图的feature数量信息
imgdir = textread('imgdirs.txt', '%s');
[class_number,~] = size(imgdir);
train_number = 15;
[sift_f_mat, sift_d_mat, feature_num] = do_sifts_2(class_number, train_number, imgdir);

%------------Kmeans----------------------
K = 128;
%do K-Means
[Idx, C] = kmeans(single(sift_d_mat'),K);

%-----------naive level 0 and level 1------
level = 1;
bin = 32;
counter_l = 1;
counter_r = 0;
imgs2cluster = [];
for i=1:class_number
    for j=1:train_number
        number = (i-1)*train_number+j;
        counter_r = counter_l + feature_num(number);
        [countM, ~] = hist(Idx(counter_l:counter_r-1), 1:K); 
        left_upper = [];left_down = []; right_upper=[]; right_down = [];
        for k=counter_l:counter_r-1
            x = sift_f_mat(1,k);
            y = sift_f_mat(2,k);
            if x <= bin && y <= bin
                left_upper = [left_upper, Idx(k)];
            elseif x<=bin && y >bin
                left_down = [left_down, Idx(k)];
            elseif x> bin && y<= bin
                right_upper = [right_upper, Idx(k)];
            else
                right_down = [right_down, Idx(k) ];
            end
        end
        counter_l = counter_r;      
        [hist_m1, ~] = hist(left_upper, 1:K);
        [hist_m2, ~] = hist(left_down, 1:K);
        [hist_m3, ~] = hist(right_upper, 1:K);
        [hist_m4, ~] = hist(right_down, 1:K);
        imgs2cluster(number, :) = [countM , hist_m1, hist_m2, hist_m3,hist_m4];
    end
end
train_lable_vector = repmat([1:class_number],train_number,1);
train_lable_vector = reshape(train_lable_vector,[],1);
train = svmtrain(train_lable_vector, imgs2cluster);

%--------------------load test image--------------
test_number = 15;
[sift_f_mat_t, sift_d_mat_t, feature_num_t] = do_sifts2_2(class_number, train_number,test_number, imgdir);

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
%-------------imgs2cluster-------------
counter_l = 1;
counter_r = 0;
imgs2cluster_t = [];
for i=1:class_number*test_number
    counter_r = counter_l + feature_num_t(i);
    [hist_m, ~] = hist(imgs_features2cluster(counter_l:counter_r-1), 1:K);
    left_upper = [];left_down = []; right_upper=[]; right_down = [];
    for k=counter_l:counter_r-1
        x = sift_f_mat(1,k);
        y = sift_f_mat(2,k);
        if x <= bin && y <= bin
            left_upper = [left_upper, Idx(k)];
        elseif x<=bin && y >bin
            left_down = [left_down, Idx(k)];
        elseif x> bin && y<= bin
            right_upper = [right_upper, Idx(k)];
        else
            right_down = [right_down, Idx(k) ];
        end
    end
    counter_l = counter_r;
    [hist_m1, ~] = hist(left_upper, 1:K);
    [hist_m2, ~] = hist(left_down, 1:K);
    [hist_m3, ~] = hist(right_upper, 1:K);
    [hist_m4, ~] = hist(right_down, 1:K);
    imgs2cluster_t(i, :) = [hist_m , hist_m1, hist_m2, hist_m3,hist_m4];
end

test_lable_vector = repmat([1:class_number],test_number,1);
test_lable_vector = reshape(test_lable_vector,[],1);
predict = svmpredict(test_lable_vector, imgs2cluster_t, train);

