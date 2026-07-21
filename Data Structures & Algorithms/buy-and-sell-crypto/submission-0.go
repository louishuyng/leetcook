func maxProfit(prices []int) int {
    sell := prices[0]
    max := 0

    for _, price := range prices {
        if price - sell > max {
            max = price - sell
        }

        if price < sell {
            sell = price
        }
    }


    return max
}
