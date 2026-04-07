public class Solution
{
    public bool IsPalindrome(int x)
    {
        string original = x.ToString();
        string reversed = new string(original.Reverse().ToArray());
        return original == reversed;
    }
}