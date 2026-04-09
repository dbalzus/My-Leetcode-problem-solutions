public class Solution
{
    public int XorAfterQueries(int[] nums, int[][] queries)
    {
        int n = nums.Length;
        foreach (var q in queries)
        {

            for (int i = q[0]; i <= q[1]; i += q[2])
            {
                nums[i] = (int)((long)nums[i] * q[3] % 1000000007);
            }
        }


        int current = 0;
        foreach (int value in nums)
        {
            current ^= value;
        }
        return current;
    }
}