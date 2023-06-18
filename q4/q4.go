package q4

//Dado um array não vazio de números inteiros "nums", cada elemento aparece duas vezes, exceto um. Encontre esse único
//elemento.
//
//Você deve implementar uma solução com complexidade de tempo linear e sem memória extra.

func SingleNumber(nums []int) int {
	numero := 0
	for i:=0;i<len(nums)-1;i++{
		for p:=i+1;p<len(nums);p++{
			if nums[i] == nums[p]{
				break
			}
		}
		numero = i
	}
	return numero
}
