//PREPEND BEGIN
#include <iostream>
#include <string>
//PREPEND END

//TEMPLATE BEGIN
//You just need to implement these two classes
#define LOG 
class Gregorian{
public:
    int tdays;   
    int MAX_DAY;
    int MIN_DAY;

    typedef struct Date{
        int year;
        std::string month;
        int day;
        Date(){
            year = 0;
        }
        Date(int y, std::string m, int d){
            year = y;
            month = m;
            day = d;
        }
        void print(){
            std::cout << "Date: "<<year <<" "<<month<<" "<<day<<std::endl;
        }
    }Date;
    
private: 
    // common 
    int days_round;
    // custom
    int days_month[12] = {31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
    std::string month[12] = {"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"};
    
   
    int month2idx(std::string m){
        for (int i = 0; i < 12; i++){
            if (m == month[i]){
                return i;
            }
        }
        return -1;
    }
    int daysFeb(int year){
        int leap = isLeap(year);
        return 29 * leap + (1-leap) * 28;
    }
    bool isLeap(int year){
        if ( year % 4000 == 0){
            return false;
        }
        if ( year % 400 == 0){
            return true;
        }
        if ( year % 100 == 0) {
            return false;
        }
        if ( year % 4 == 0) {
            return true;
        }else {
            return false;
        }
    }
    int daysYear(int year){
        bool leap = isLeap(year);
        return 366 * leap + (1-leap) * 365;
    }
    int date2days(Date d){
        int year = d.year;
        int month_idx = month2idx(d.month);
        int day = d.day;
        days_month[1] = daysFeb(year);
        for (int i = 0; i < month_idx; i++){
            day += days_month[i];
        }
        year--;
        int leap_days = year/4 - year/100 + year/400 - year/4000;
        day = day + leap_days + year * 365;
        return day;
    }
    Date days2date(int t){
        int tday = t;
        if ( tday > MAX_DAY || tday < MIN_DAY ){
            return Date();
        }
        int year = 1;
        while ( tday > days_round ){
            year += 4000;
            tday -= days_round;
        }
        int days_of_year = daysYear(year);
        while (tday > days_of_year){
            tday -= days_of_year;
            year++;
            days_of_year = daysYear(year);
        }
        int month_idx = 0;
        days_month[1] = daysFeb(year);
        for (int i = 0; i < 12; i++){
            if ( tday > days_month[i]){
                tday -= days_month[i];
            }else{
                Date d = Date(year, month[i], tday);
                return d;
            }
        }
    }
public:
    Gregorian(){
        days_round = 4000 * 365 + 4000/4 - 4000/100 + 4000/400 - 4000/4000;
        MAX_DAY = 999999 * 365 + 999999/4 - 999999/100 + 999999/400 - 999999/4000;
        MIN_DAY = 1;
    }
    Gregorian(int year, char* month, int day){
        days_round = 4000 * 365 + 4000/4 - 4000/100 + 4000/400 - 4000/4000;
        MAX_DAY = 999999 * 365 + 999999/4 - 999999/100 + 999999/400 - 999999/4000;
        MIN_DAY = 1;
        tdays = date2days(Date(year, std::string(month), day));
    }
    void print_today() {
        Date d = days2date(tdays);
        std::cout << d.year <<" " << d.month << " " << d.day << std::endl;
    }
    void print_month() {
        std::string top    = std::string("┌────┬────┬────┬────┬────┬────┬────┐");
        
        std::string title  = std::string("│ Sun│ Mon│Tues│ Wed│Thur│ Fri│ Sat│");
        
        std::string mid    = std::string("├────┼────┼────┼────┼────┼────┼────┤");
        
        std::string bottom = std::string("└────┴────┴────┴────┴────┴────┴────┘");
        Date d = days2date(tdays);
        int w = (tdays - d.day + 1) % 7;
        int m = month2idx(d.month);
        std::string pad[42];
        days_month[1] = daysFeb(d.year);
        for (int i = 0; i < 42; i++) { pad[i] = "  ";}
        int j = 1;
        for (int i = w; i < w+days_month[(m)%12]; i++){
            std::string num = std::to_string(j);
            if (j < 10){
                num = std::string("0") + num;
            }
            pad[i] =  num ;
            j++;
        }
        std::cout << d.month <<"                                 " << std::endl;
        std::cout << top << "\n" << title << "\n" << mid << std::endl;
        for (int i = 0; i < 6;  i++){
            std::cout << "│ " << pad[i*7] <<" │ " << pad[i*7+1] << " │ " << pad[i*7+2] << " │ " << pad[i*7+3] <<" │ " << pad[i*7+4] <<" │ " << pad[i*7+5] <<" │ " << pad[i*7+6] <<" │" << std::endl; 
            if ( i == 5 ) {
                std::cout << bottom << std::endl;
            }else{
                std::cout << mid << std::endl;
            }
        }

    }
    void print_year(){
        std::string top    = std::string("┌────┬────┬────┬────┬────┬────┬────┐") ;
        top = top + " " + top + " " + top;
        std::string title  = std::string("│ Sun│ Mon│Tues│ Wed│Thur│ Fri│ Sat│");
        title = title + " " + title + " " + title;
        std::string mid    = std::string("├────┼────┼────┼────┼────┼────┼────┤");
        mid = mid + " " + mid + " " + mid;
        std::string bottom = std::string("└────┴────┴────┴────┴────┴────┴────┘");
        bottom = bottom + " " + bottom + " " + bottom;

        Date d = days2date(tdays);
        int tday = date2days(Date(d.year, month[0].c_str(),1));
        int w = tday % 7;
        std::string pad[12][42];
        days_month[1] = daysFeb(d.year);
        for (int i = 0; i < 12; i++) { 
            for (int j = 0; j < 42; j++){
                pad[i][j] = std::string("  ");
            }
        }
        for (int m = 0; m < 12; m++){
            int j = 1;
            for (int i = w; i < w+days_month[(m)%12]; i++){
                std::string num = std::to_string(j);
                if (j < 10){
                    num = std::string("0") + num;
                }
                pad[m][i] =  num ;
                j++;
            }
            w = (w + days_month[m]) % 7;
        }
        for (int i = 0; i < 4; i++){
            for (int j = 0; j < 3; j++){
                std::cout << month[i*3+j] <<"                                 ";
                if ( j < 2 ){
                    std::cout << " ";
                }
            }
            std::cout << std::endl;
            std::cout << top << "\n" << title << "\n" << mid << std::endl;
            for (int k = 0; k < 6;  k++){
                for (int j = 0; j < 3; j++){
                    std::cout << "│ " << pad[i*3+j][k*7] << \
                                " │ " << pad[i*3+j][k*7+1] << \
                                " │ " << pad[i*3+j][k*7+2] << \
                                " │ " << pad[i*3+j][k*7+3] << \
                                " │ " << pad[i*3+j][k*7+4] << \
                                " │ " << pad[i*3+j][k*7+5] << \
                                " │ " << pad[i*3+j][k*7+6] << \
                                " │" ;
                    if ( j < 2){
                        std::cout << " ";
                    } 
                }
                std::cout << std::endl;
                if ( k == 5 ) {
                    std::cout << bottom << std::endl;
                }else{
                    std::cout << mid << std::endl;
                }
            }
        }
    }
    bool go_to(int year, const char* month, int day){
        if (year < 1 || day < 1 || month2idx(std::string(month))==-1){
            return false;
        }
        Date d = Date(year, std::string(month), day);
        int t = date2days(Date(year, std::string(month), day));
        if (t > MAX_DAY || t < MIN_DAY){
            return false;
        }else{
            tdays = t;
            return true;
        }
    }
    bool pass_day(int num_days){
        #ifdef LOG
            std::cout<<"--- G S pass_day : from "<<tdays<<" days + "<<num_days<<" to ";
        #endif      
        int t = tdays + num_days;
        if (t > MAX_DAY || t < MIN_DAY){
            #ifdef LOG
                std::cout << tdays<<" days"<<std::endl;
            #endif
            return false;
        }
        tdays = t;
        #ifdef LOG
            std::cout << tdays<<" days"<<std::endl;
        #endif
        return true;       
    }
    bool pass_month(int num_months){
        Date d = days2date(tdays);
        int m = (d.year -1) * 12 + month2idx(d.month) + 1;
        m += num_months;
        #ifdef LOG
            std::cout << "--- G pass_month : from "<<d.year<<" "<<d.month<<" "<<d.day<<" + "<<num_months<<" months to ";
        #endif
        int year = (m / 12);
        m= m % 12 ;
        Date c = m < 0 ? Date(-1, " ", -1) : m == 0 ? Date(year, "Dec", 1) : Date(year + 1, month[m-1].c_str(), 1);
        bool f = go_to(c.year, c.month.c_str(), 1);
        
        #ifdef LOG
            Date t = days2date(tdays);
            t.print();
        #endif
        return f;
    }
    bool pass_year(int num_years){
        Date d = days2date(tdays);
        int year = d.year;
        year += num_years;

        #ifdef LOG
            std::cout << "--- G pass_year : from "<<d.year<<" "<<d.month<<" "<<d.day<<" + "<<num_years<<" years to ";
        #endif
        bool f = go_to(year, month[0].c_str(), 1);
        #ifdef LOG
            Date t = days2date(tdays);
            t.print();
        #endif
        return f;
    }

};

class Shanghaitech:public Gregorian{
private:
    std::string hex[16] = {"0","1","2","3","4","5","6","7","8","9","A","B","C","D","E","F"};
    
    std::string int2hex(int i){
        if (i < 16){
            return "0"+hex[i];
        }else{
            return hex[i/16] + hex[i%16];
        }
    }
    std::string month[9] = {"Sist", "Spst", "Slst", "Sem", "Sca", "Ims", "Ihuman", "Siais", "Ih"};
    std::string leap_year[9][10] = {
        {"Sist", "SIST","Spst", "Slst", "Sem", "Sca", "Ims", "Ihuman", "Siais", "Ih"},
        {"Sist", "Spst", "SPST", "Slst", "Sem", "Sca", "Ims", "Ihuman", "Siais", "Ih"},
        {"Sist", "Spst", "Slst", "SLST","Sem", "Sca", "Ims", "Ihuman", "Siais", "Ih"},
        {"Sist", "Spst", "Slst", "Sem", "SEM","Sca", "Ims", "Ihuman", "Siais", "Ih"},
        {"Sist", "Spst", "Slst", "Sem", "Sca", "SCA","Ims", "Ihuman", "Siais", "Ih"},
        {"Sist", "Spst", "Slst", "Sem", "Sca", "Ims", "IMS", "Ihuman", "Siais", "Ih"},
        {"Sist", "Spst", "Slst", "Sem", "Sca", "Ims", "Ihuman", "IHUMAN", "Siais", "Ih"},
        {"Sist", "Spst", "Slst", "Sem", "Sca", "Ims", "Ihuman", "Siais", "SIAIS","Ih"},
        {"Sist", "Spst", "Slst", "Sem", "Sca", "Ims", "Ihuman", "Siais", "Ih", "IH"}
    };
    std::string leap_month[9] = {"SIST", "SPST", "SLST", "SEM", "SCA", "IMS", "IHUMAN", "SIAIS", "IH"};
    int days_month[3][9] = {
        {41,40,39,41,40,39,41,40,39},
        {39,41,40,39,41,40,39,41,40},
        {40,39,41,40,39,41,40,39,41},
            
    };
    int days_year = 360;

    int years_table = 4000;
    int days_table[253] = {
    1459166, 1460780, 1460777, 1461956, 1460690, 1461962, 1461836, 1462107, 1460358, 1461974,
    1461909, 1462118, 1461471, 1462158, 1461754, 1461347, 1459691, 1461839, 1461869, 1462176,
    1461435, 1462303, 1461793, 1461516, 1460493, 1462188, 1461830, 1461675, 1460686, 1461715,
    1460608, 1460318, 1458605, 1461306, 1461869, 1462140, 1461471, 1462149, 1461790, 1461817,
    1460403, 1461993, 1461867, 1461950, 1460722, 1461835, 1460799, 1460647, 1459232, 1461219,
    1462147, 1461714, 1461072, 1461444, 1461190, 1460647, 1459579, 1460635, 1461386, 1460647,
    1459813, 1459877, 1459891, 1459003, 1458425, 1458719, 1461568, 1461872, 1462230, 1461390,
    1462197, 1461790, 1461718, 1460298, 1462288, 1461714, 1461805, 1460593, 1461808, 1460644,
    1460675, 1458702, 1462113, 1461745, 1461826, 1460721, 1461910, 1460763, 1460723, 1459156,
    1461605, 1460995, 1460725, 1459230, 1460538, 1459308, 1459158, 1458120, 1460517, 1462200,
    1461793, 1461426, 1460933, 1461428, 1460608, 1459851, 1460267, 1461557, 1460725, 1459928,
    1459840, 1460006, 1459198, 1458460, 1459121, 1461757, 1460647, 1460356, 1459449, 1460397,
    1459200, 1458807, 1458343, 1460556, 1459043, 1459119, 1457941, 1459079, 1458009, 1458132,
    1457940, 1460759, 1461959, 1461833, 1462147, 1461479, 1462071, 1461793, 1461110, 1461044,
    1462158, 1461793, 1461387, 1461011, 1461468, 1460568, 1459812, 1459888, 1462179, 1461908,
    1461516, 1460845, 1461713, 1460568, 1460084, 1459723, 1461715, 1460647, 1460358, 1459329,
    1460437, 1459200, 1458650, 1458171, 1462143, 1461790, 1461817, 1460572, 1461865, 1460602,
    1460680, 1459157, 1461913, 1460721, 1460764, 1459196, 1460649, 1459232, 1459159, 1458045,
    1461191, 1461149, 1460764, 1459503, 1460344, 1459579, 1459041, 1458195, 1459934, 1459891,
    1459082, 1458231, 1458865, 1458425, 1458133, 1457715, 1459584, 1462233, 1461754, 1461754,
    1460683, 1461718, 1460683, 1460438, 1459240, 1461805, 1460608, 1460554, 1459133, 1460635,
    1459122, 1459000, 1457939, 1461909, 1460723, 1460804, 1459116, 1460807, 1459155, 1459158,
    1458006, 1460742, 1459308, 1459082, 1458120, 1459171, 1458120, 1458093, 1457727, 1460952,
    1461468, 1460568, 1459891, 1459799, 1459890, 1459158, 1458385, 1459665, 1460006, 1459082,
    1458460, 1458873, 1458419, 1458174, 1457711, 1459287, 1460397, 1459082, 1458807, 1458331,
    1458845, 1458174, 1457861};

    int months_table[253] = {
    36478, 36518, 36518, 36549, 36516, 36549, 36546, 36555, 36508, 36549,
    36548, 36555, 36537, 36556, 36546, 36537, 36492, 36545, 36547, 36556,
    36537, 36559, 36547, 36541, 36513, 36556, 36548, 36545, 36519, 36546,
    36518, 36511, 36466, 36531, 36547, 36554, 36539, 36554, 36547, 36548,
    36513, 36550, 36549, 36551, 36521, 36548, 36523, 36519, 36483, 36530,
    36556, 36545, 36530, 36538, 36533, 36519, 36492, 36517, 36538, 36519,
    36498, 36499, 36500, 36477, 36462, 36467, 36538, 36547, 36557, 36536,
    36556, 36547, 36546, 36509, 36558, 36545, 36548, 36517, 36548, 36519,
    36520, 36469, 36553, 36546, 36548, 36521, 36550, 36522, 36521, 36481,
    36542, 36528, 36521, 36483, 36516, 36485, 36481, 36454, 36512, 36557,
    36547, 36539, 36525, 36539, 36518, 36499, 36508, 36542, 36521, 36501,
    36498, 36503, 36482, 36463, 36478, 36547, 36519, 36512, 36488, 36513,
    36482, 36472, 36459, 36517, 36478, 36480, 36449, 36479, 36451, 36454,
    36448, 36517, 36549, 36546, 36556, 36537, 36554, 36547, 36531, 36526,
    36556, 36547, 36538, 36527, 36540, 36517, 36498, 36497, 36556, 36550,
    36541, 36523, 36546, 36517, 36505, 36494, 36546, 36519, 36512, 36485,
    36514, 36482, 36468, 36454, 36554, 36547, 36548, 36517, 36549, 36518,
    36520, 36481, 36550, 36521, 36522, 36482, 36519, 36483, 36481, 36452,
    36531, 36532, 36522, 36490, 36511, 36492, 36478, 36456, 36500, 36500,
    36479, 36457, 36473, 36462, 36454, 36443, 36488, 36557, 36546, 36547,
    36519, 36546, 36520, 36514, 36482, 36548, 36518, 36517, 36480, 36519,
    36480, 36477, 36449, 36550, 36521, 36523, 36480, 36523, 36481, 36481,
    36451, 36521, 36485, 36479, 36454, 36481, 36454, 36453, 36443, 36524,
    36540, 36517, 36500, 36497, 36500, 36481, 36461, 36493, 36503, 36479,
    36463, 36473, 36462, 36455, 36443, 36482, 36513, 36479, 36472, 36459,
    36473, 36455, 36447};

    int month2idx(std::string s){
        for (int i = 0; i < 9; i++){
            if (s == month[i]){
                return i;
            }
        }
        return -1;
    }
    int leapMonth2idx(std::string s){
        for (int i = 0; i < 9; i++){
            if (s == leap_month[i]){
                return i;
            }
        }
        return -1;
    }
    int daysMonth(int year, std::string month){
        int month_idx = month2idx(month);
        if (month_idx == -1){
            month_idx = leapMonth2idx(month);
        }
        int days = days_month[year%3][month_idx];
        return days;
    }
    int daysYear(int year){
        bool leap = isLeap(year);
        int days = days_year;
        if (leap){
            days += days_month[year%3][leapMonth(year)-1];
        }
        return days;
    }
    int monthsYear(int year){
        bool leap = isLeap(year);
        return 10 * leap + (1-leap)*9;
    }
    int bitCount(int n){
        unsigned int count = 0;
        while (n) {
            count += n & 1;
            n >>= 1;
        }
        return count;
    }
    bool isLeap(int year){
        int count = bitCount(year);
        int mod = (year + count)%8;
        _GLIBCXX_DEBUG_ASSERT(mod >= 0);
        if (mod == 0){
            return true;
        } else {
            return false;
        }
    }
    int leapMonth(int year){
        int count = bitCount(year);
        int mod = (year - count) % 9  + 1;
        _GLIBCXX_DEBUG_ASSERT(mod >= 0);
        return mod;
    }
    int days2months(int t){
        int months = 1;
        int year = 1;
        for (int i = 0; i < sizeof(months_table)/sizeof(months_table[0]);i++){
            if (t > days_table[i] ){
                months += months_table[i];
                year += years_table;
                t -= days_table[i];
            }else{
                break;
            }
        }
        int day_of_year = daysYear(year);
        while (t > day_of_year){
            months += monthsYear(year);
            year++;
            t -= day_of_year;
            day_of_year = daysYear(year);
        }
        bool leap = isLeap(year);
        bool leap_month_idx = leapMonth(year)-1;
        if (leap){
            for (int i = 0; i < 10;i++){
                int day_of_month = daysMonth(year, leap_year[leap_month_idx][i]);
                if (t > day_of_month){
                    months++;
                    t -= day_of_month;
                }else{
                    break;
                }
            }
        }else{
            for (int i = 0; i < 9;i++){
                int day_of_month = daysMonth(year, month[i]);
                if (t > day_of_month){
                    months++;
                    t -= day_of_month;
                }else{
                    break;
                }
            }
        }
        return months;
    }
    int months2days(int m){
        if (m < 1){
            return -1;
        }
        int d = 1;
        int year = 1;
        for (int i = 0; i < sizeof(months_table)/sizeof(months_table[0]);i++){
            if (m > months_table[i] ){
                d += days_table[i];
                year += years_table;
                m -= months_table[i];
            }else{
                break;
            }
        }
        int month_of_year = monthsYear(year);
        while (m > month_of_year){
            m -= month_of_year;
            d += daysYear(year);
            year++;
            month_of_year = monthsYear(year);
        }
        bool leap = isLeap(year);
        bool leap_month_idx = leapMonth(year)-1;
        for (int i = 0; i < m - 1; i++){
            if (leap){
                int day_of_month = daysMonth(year, leap_year[leap_month_idx][i]);
                d += day_of_month;
            }else{
                int day_of_month = daysMonth(year, month[i]);
                d += day_of_month;
            }
        }
        return d;
    }
    Date days2date(int t){
        int days = t;
        int year = 1;
        for (int i = 0; i < sizeof(days_table)/sizeof(days_table[0]); i++){
            if (days > days_table[i]){
                year += years_table;
                days -= days_table[i];
            }else{
                break;
            }
        }
        int days_year = daysYear(year);
        while (days > days_year){
            year++;
            days -= days_year;
            days_year = daysYear(year);
        }
        std::string m = " ";
        if (isLeap(year)){
            int leap_month_idx = leapMonth(year) - 1;
            for (int i = 0; i < 10; i++){
                int tmp = daysMonth(year, leap_year[leap_month_idx][i]); 
                if ( days > tmp){
                    days -= tmp;
                }else{
                    m = leap_year[leap_month_idx][i];
                    break;
                }
            }
        }else{
            for (int i = 0; i < 9; i++){
                int tmp = daysMonth(year, month[i]); 
                if ( days > tmp){
                    days -= tmp;
                }else{
                    m = month[i];
                    break;
                }
            }
        }
        int day = days;
        return Date(year, m, day);
    }
    int date2days(Date d){
        int year = 1;
        int days = 0;
        for (int i = 0; i < sizeof(days_table)/sizeof(int); i++){
            if (year + years_table < d.year){
                year += years_table;
                days += days_table[i];
            }else{
                break;
            }
        }
        while ( year < d.year){
            days += daysYear(year);
            year++;
        }
        if (isLeap(year)){
            int leap_month_idx = leapMonth(year) - 1;
            //std::cout << "------ leap_month_idx : " << leap_month_idx<<" of "<<year <<" year"<<std::endl; 
            int target_month_idx = -1;
            for (int i = 0; i < 10; i++){
                if (d.month == leap_year[leap_month_idx][i]){
                    target_month_idx = i;
                    break;
                }
            }
            if (target_month_idx == -1){
                return tdays;
            }
            for (int i = 0; i < target_month_idx; i++){
                days += daysMonth(year, leap_year[leap_month_idx][i]);
            }
        }else{
            int target_month_idx = -1;
            for (int i = 0; i < 10; i++){
                if (d.month == month[i]){
                    target_month_idx = i;
                    break;
                }
            }
            if (target_month_idx == -1){
                return tdays;
            }
            for (int i = 0; i < target_month_idx; i++){
                days += daysMonth(year, month[i]);
            }
        }
        days += d.day;
        //std::cout<<"---date2days---S]"<<std::endl;
        return days;
    }
public:
    Shanghaitech(int year, char* month, int day):Gregorian(){
        tdays = date2days(Date(year, std::string(month),day));
        #ifdef LOG
            std::cout<<"Init MAX_DAY: "<<MAX_DAY<<" MIN_DAY: "<<MIN_DAY<<std::endl;
            std::cout<<"Init Shanghaitech with "<<year<<" " <<month<<" "<<day<<" as "<< tdays <<" days"<<std::endl;
            Date d = days2date(tdays);
            std::cout << "Init check: " << tdays << " as ";d.print();
        #endif
    }
    void print_today() {
        //std::cout<<"[S---print_today--- "<<std::endl;
        Date d = days2date(tdays);
        std::cout << d.year <<" " << d.month << " " << d.day << std::endl;
        //std::cout<<"---print_today---S] "<<std::endl;
    };
    void print_month() {
        std::string top    = std::string("┌────┬────┬────┬────┬────┬────┬────┐");
        
        std::string title  = std::string("│ Sun│ Mon│Tues│ Wed│Thur│ Fri│ Sat│");
        
        std::string mid    = std::string("├────┼────┼────┼────┼────┼────┼────┤");
        
        std::string bottom = std::string("└────┴────┴────┴────┴────┴────┴────┘");

        //std::cout<<"[S---print_month--- "<<std::endl;

        Date d = days2date(tdays);
        int w = (tdays - d.day + 1) % 7;
        int m = month2idx(d.month);
        int days = daysMonth(d.year, d.month);
        std::string pad[49];
        for (int i = 0; i < 49; i++) { pad[i] = "  ";}
        int j = 1;
        for (int i = w; i < w+days; i++){
            pad[i] =  int2hex(j);
            j++;
        }
        std::string name = d.month + std::string("                                     ");
        name.resize(36);
        std::cout << name << std::endl;
        std::cout << top << "\n" << title << "\n" << mid << std::endl;
        for (int i = 0; i < 7;  i++){
            std::cout << "│ " << pad[i*7] <<" │ " << pad[i*7+1] << " │ " << pad[i*7+2] << " │ " << pad[i*7+3] <<" │ " << pad[i*7+4] <<" │ " << pad[i*7+5] <<" │ " << pad[i*7+6] <<" │" << std::endl; 
            if ( i == 6 ) {
                std::cout << bottom << std::endl;
            }else{
                std::cout << mid << std::endl;
            }
        }
        //std::cout<<"---print_month---S] "<<std::endl;

    }
    void print_year(){
        std::string top    = std::string("┌────┬────┬────┬────┬────┬────┬────┐") ;
        top = top + " " + top + " " + top;
        std::string title  = std::string("│ Sun│ Mon│Tues│ Wed│Thur│ Fri│ Sat│");
        title = title + " " + title + " " + title;
        std::string mid    = std::string("├────┼────┼────┼────┼────┼────┼────┤");
        mid = mid + " " + mid + " " + mid;
        std::string bottom = std::string("└────┴────┴────┴────┴────┴────┴────┘");
        bottom = bottom + " " + bottom + " " + bottom;

        //std::cout<<"[S---print_year--- "<<std::endl;
        Date d = days2date(tdays);

        int tday = date2days(Date(d.year, month[0], 1));
        int w = tday % 7;
        std::string pad[10][49];
        for (int i = 0; i < 10; i++) { 
            for (int j = 0; j < 49; j++){
                pad[i][j] = std::string("  ");
            }
        }
        int days[10];
        bool leap = isLeap(d.year);
        int leap_month_idx = leapMonth(d.year)-1;
        if (isLeap(d.year)){
            for (int i = 0; i < 10; i++){
                days[i] = daysMonth(d.year, leap_year[leap_month_idx][i]);
            }
        }else{
            for (int i = 0; i < 9; i++){
                days[i] = daysMonth(d.year, month[i]);
            }
            days[9] = 0;
        }
        for (int m = 0; m < 10; m++){
            int j = 1;
            for (int i = w; i < w+days[m]; i++){
                pad[m][i] =  int2hex(j) ;
                j++;
            }
            w = (w + days[m]) % 7;
        }

        for (int i = 0; i < 3; i++){
            for (int j = 0; j < 3; j++){
                if (leap){
                    std::string name = leap_year[leap_month_idx][i*3+j] + std::string("                                   ");
                    name.resize(36);
                    std::cout << name;
                }else{
                    std::string name = month[i*3+j] + std::string("                                   ");
                    name.resize(36);
                    std::cout << name;
                }
                if ( j < 2 ){
                    std::cout << " ";
                }
            }
            std::cout << std::endl;
            std::cout << top << "\n" << title << "\n" << mid << std::endl;
            for (int k = 0; k < 7;  k++){
                for (int j = 0; j < 3; j++){
                    std::cout << "│ " << pad[i*3+j][k*7] << \
                                " │ " << pad[i*3+j][k*7+1] << \
                                " │ " << pad[i*3+j][k*7+2] << \
                                " │ " << pad[i*3+j][k*7+3] << \
                                " │ " << pad[i*3+j][k*7+4] << \
                                " │ " << pad[i*3+j][k*7+5] << \
                                " │ " << pad[i*3+j][k*7+6] << \
                                " │" ;
                    if ( j < 2){
                        std::cout << " ";
                    } 
                }
                std::cout << std::endl;
                if ( k == 6 ) {
                    std::cout << bottom << std::endl;
                }else{
                    std::cout << mid << std::endl;
                }
            }
        }
        if (leap){
            std::string etop    = std::string("┌────┬────┬────┬────┬────┬────┬────┐");
            
            std::string etitle  = std::string("│ Sun│ Mon│Tues│ Wed│Thur│ Fri│ Sat│");
            
            std::string emid    = std::string("├────┼────┼────┼────┼────┼────┼────┤");
            
            std::string ebottom = std::string("└────┴────┴────┴────┴────┴────┴────┘");
            std::string name = leap_year[leap_month_idx][9] + std::string("                                   ");
            name.resize(36);
            std::cout << name << std::endl;
            std::cout << etop << "\n" << etitle << "\n" << emid << std::endl;
            for (int i = 0; i < 7;  i++){
                std::cout << "│ " << pad[9][i*7] <<" │ " << pad[9][i*7+1] << " │ " << pad[9][i*7+2] << " │ " << pad[9][i*7+3] <<" │ " << pad[9][i*7+4] <<" │ " << pad[9][i*7+5] <<" │ " << pad[9][i*7+6] <<" │" << std::endl; 
                if ( i == 6 ) {
                    std::cout << ebottom << std::endl;
                }else{
                    std::cout << emid << std::endl;
                }
            }
        }
        //std::cout<<"---print_year---S] "<<std::endl;
    }
    bool go_to(int year, const char* month, int day){
        Date d = Date(year, std::string(month), day);
        int t = date2days(d);
        //std::cout << "--- go_to " << d.year << " " << d.month << " " << d.day << " as " << t << " days"<<std::endl;

        //Date c = days2date(t);
        //std::cout << "--- go_to check: " << t << " days as "; c.print();
        if (t > MAX_DAY || t < MIN_DAY){
            return false;
        }else{
            tdays = t;
            return true;
        }
    }
    bool pass_day(int num_days){
        return Gregorian::pass_day(num_days);       
    }
    bool pass_month(int num_months){
        int m = days2months(tdays);
        m += num_months;

        //std::cout << days2months(tdays) << " "<< days2months(months2days(days2months(tdays))) << std::endl;
        #ifdef LOG
            Date d = days2date(tdays);
            std::cout << "--- S pass_month : from "<<d.year<<" "<<d.month<<" "<<d.day<<" + "<<num_months<<" months to ";
        #endif

        int t = months2days(m);
        if (t > MAX_DAY || t < MIN_DAY){
            #ifdef LOG
                Date d = days2date(tdays);
                d.print();
            #endif
            return false;
        }
        tdays = t;
        Date c = days2date(tdays);
        #ifdef LOG
            c.print();
        #endif
        return true;
    }
    bool pass_year(int num_years){
        //std::cout<<"[S---pass_year--- "<< num_years<<std::endl;
        Date d = days2date(tdays);
        #ifdef LOG
            std::cout << "--- S pass_year : from " << d.year <<" " << d.month << " " << d.day <<" + "<< num_years<< " years to ";
        #endif

        if (d.year + num_years < 1){
            #ifdef LOG
                Date e = days2date(tdays);
                e.print();
            #endif
            
            return false;
        }

        bool f = go_to(d.year + num_years, month[0].c_str(), 1);
        #ifdef LOG
            Date e = days2date(tdays);
            e.print();
        #endif
        return f;
    }
};

//TEMPLATE END

//APPEND BEGIN

// Use this main function
int main()
{
    int year, day, n;
    std::string calendar, f;
    char month[10];

    std::cin >> f;
    std::cin >> year >> month >> day;
    
    Shanghaitech date(year, month, day);
    Shanghaitech* S = &date;
    Gregorian* G = &date;
    std::cin >> calendar;
    std::cin >> f;
    while (!std::cin.eof()){
        if (f == "pass_day"){
            std::cin >> n;
            if (calendar=="G") G->pass_day(n); else S->pass_day(n);
        }
        else if (f == "pass_month"){
            std::cin >> n;
            if (calendar=="G") G->pass_month(n); else S->pass_month(n);
        }
        else if (f == "pass_year"){
            std::cin >> n;
            if (calendar=="G") G->pass_year(n); else S->pass_year(n);
        }
        else if (f == "print_today"){
            if (calendar=="G") G->print_today(); else S->print_today();
        }
        else if (f == "print_month"){
            if (calendar=="G") G->print_month(); else S->print_month();
        }
        else if (f == "print_year"){
            if (calendar=="G") G->print_year(); else S->print_year();
        }
        else if (f == "go_to"){
            std::cin >> year >> month >> day;
            if (calendar=="G") G->go_to(year, month, day); else S->go_to(year, month, day);
        }
        calendar = "";
        f = "";
        std::cin >> calendar;
        std::cin >> f;
    }
    return 0;
}
// int main()
// {
//     int year, day, n;
//     std::string calendar, f;
//     char month[12];

//     std::cin >> calendar;
//     std::cin >> year >> month >> day;
//     Gregorian date(year, month, day);
//     Gregorian* G = &date;
//     std::cin >> f;
//     while (!std::cin.eof()){
//         if (f == "pass_day"){
//             std::cin >> n;
//              G->pass_day(n); 
//         }
//         else if (f == "pass_month"){
//             std::cin >> n;
//             G->pass_month(n); 
//         }
//         else if (f == "pass_year"){
//             std::cin >> n;
//             G->pass_year(n); 
//         }
//         else if (f == "print_today"){
//              G->print_today(); 
//         }
//         else if (f == "print_month"){
//             G->print_month(); 
//         }
//         else if (f == "print_year"){
//              G->print_year();
//         }
//         else if (f == "go_to"){
//             std::cin >> year >> month >> day;
//             G->go_to(year, month, day);
//         }
//         f = "";
//         std::cin >> f;
//     }
//     return 0;
// }
//APPEND END