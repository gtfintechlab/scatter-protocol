// SPDX-License-Identifier: MIT
pragma solidity >=0.8.17;

library ScatterMath {
    function floorSqrt(uint256 n) internal pure returns (uint256) {
        unchecked {
            if (n > 0) {
                uint256 x = n / 2 + 1;
                uint256 y = (x + n / x) / 2;
                while (x > y) {
                    x = y;
                    y = (x + n / x) / 2;
                }
                return x;
            }
            return 0;
        }
    }
}
