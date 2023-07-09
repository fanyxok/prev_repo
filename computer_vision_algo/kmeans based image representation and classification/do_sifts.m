function [sift_d_mat,feature_num] = do_sifts(class_number,train_number,imgdir)

sift_d_mat = [];
feature_num = [];
for i=1:class_number
    i
    for j=1:train_number
        if i<10 
            if j < 10
                img_id = ['\00' num2str(i) '_' '000' num2str(j) '.jpg'];
            else 
                img_id = ['\00' num2str(i) '_' '00' num2str(j) '.jpg'];        
            end
        elseif i < 100
            if j < 10
                img_id = ['\0' num2str(i) '_' '000' num2str(j) '.jpg'];
            else 
                img_id = ['\0' num2str(i) '_' '00' num2str(j) '.jpg'];        
            end
        else 
            if j < 10
                img_id = ['\' num2str(i) '_' '000' num2str(j) '.jpg'];
            else 
                img_id = ['\' num2str(i) '_' '00' num2str(j) '.jpg'];        
            end
        end
        img_path = strcat(char(imgdir(i)),img_id);
        img_read = imread(img_path);
        img_read = imresize(img_read, [64,64]);
        if (size(img_read,3) == 1)
            img_gray = single(img_read);
        else
            img_gray = single(rgb2gray(img_read));
        end    
        [~, sift_d] = vl_dsift(img_gray,'size',4, 'step',4);
        [~, n] = size(sift_d);
        feature_num = [feature_num n];
        sift_d_mat = [sift_d_mat sift_d];
    end
end

end

