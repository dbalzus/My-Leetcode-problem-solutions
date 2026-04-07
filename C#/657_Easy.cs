public class Solution
{
    public bool JudgeCircle(string moves)
    {
        int up = countMoves(moves, 'U');
        int down = countMoves(moves, 'D');
        int left = countMoves(moves, 'L');
        int right = countMoves(moves, 'R');

        if ((up == down) && (left == right))
        {
            return true;
        }
        else
        {
            return false;
        }

    }

    public int countMoves(string moves, char? direction)
    {
        int count = 0;
        foreach (char c in moves)
        {
            if (c == direction) count++;
        }
        return count;
    }
}