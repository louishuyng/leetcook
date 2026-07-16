class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    canPermutePalindrome(s) {
        let count = 0;
        for (let i = 0; i < 128 && count <= 1; i++) {
            let ct = 0;
            for (let j = 0; j < s.length; j++) {
                if (s.charCodeAt(j) === i) ct++;
            }
            count += ct % 2;
        }
        return count <= 1;
    }
}