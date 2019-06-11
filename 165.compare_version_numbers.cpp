class Solution {
public:
    int compareVersion(string version1, string version2) {
    size_t i = 0;
    size_t j = 0;
    string str_1, str_2 = "0";
    while(i < version1.size() || j < version2.size()) {
        while (str_1.back() != '.' && i < version1.size()) str_1 += version1[i++];
        while (str_2.back() != '.' && j < version2.size()) str_2 += version2[j++];
        if (stoi(str_1) > stoi(str_2)) {
            return 1;
        } else if (stoi(str_1) < stoi(str_2)) {
            return -1;
        }
        str_1 = "0";
        str_2 = "0";
    }
    return 0;
    }
};
