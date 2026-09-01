func maxProfit(prices []int) int {

	if len(prices) < 2{
	return 0
	}
	min_price := prices[0]
	max_price := 0

	for i:=0;i<len(prices);i++{
		if prices[i] < min_price{
			min_price = prices[i]
		}else if prices[i] - min_price > max_price{
			max_price = prices[i] - min_price
		}
	}
	return max_price
}