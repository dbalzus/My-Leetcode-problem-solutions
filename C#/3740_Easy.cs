public class Solution
{
    public int MinimumDistance(int[] nums)
    {
        int n = nums.Length;
        int ans = n + 1;
        if (n < 3)
        {
            return -1;
        }
        for (int i = 0; i < n; i++)
        {
            for (int j = i + 1; j < n; j++)
            {
                if (nums[i] != nums[j])
                {
                    continue;
                }
                for (int k = j + 1; k < n; k++)
                {
                    if (nums[j] == nums[k])
                    {
                        ans = Math.Min(ans, k - i);
                        break;
                    }
                }
            }
        }
        if (ans == n + 1)
        {
            return -1;
        }
        else
        {
            return ans * 2;
        }
    }
}