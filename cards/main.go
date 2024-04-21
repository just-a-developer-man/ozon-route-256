package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	TheNewesSolution(in, out)
}

func Solution(in io.Reader, out io.Writer) {
	var friendsCount, cardsCount int
	fmt.Fscan(in, &friendsCount, &cardsCount)

	friendsCards := make(map[int]int)
	for friendIndex := 0; friendIndex < friendsCount; friendIndex++ {
		var cardsAmount int
		fmt.Fscan(in, &cardsAmount)
		friendsCards[friendIndex] = cardsAmount
	}

	myCards := make(map[int]struct{})
	for cardIndex := 1; cardIndex <= cardsCount; cardIndex++ {
		myCards[cardIndex] = struct{}{}
	}

	solutionSlice := make([]int, friendsCount)

	for len(friendsCards) > 0 {

		err := CalculateCard(friendsCards, myCards, &solutionSlice)
		if err != nil {
			fmt.Fprintln(out, -1)
			return
		}
	}

	for _, value := range solutionSlice {
		fmt.Fprint(out, value)
		fmt.Fprint(out, " ")
	}
	fmt.Fprintln(out)

}

func CalculateCard(friendsCards map[int]int, myCards map[int]struct{}, solutionSlice *[]int) error {
	friendIndex := FindFriendWithMaxCardIndex(friendsCards)
	myCardIndex, err := FindMyCardMinIndex(friendIndex, friendsCards, myCards)

	if err != nil {
		return err
	}

	(*solutionSlice)[friendIndex] = myCardIndex
	delete(friendsCards, friendIndex)
	delete(myCards, myCardIndex)

	return nil
}

func FindFriendWithMaxCardIndex(friendsCards map[int]int) (friendIndex int) {
	maxCardIndex := -1

	for index, friendCardIndex := range friendsCards {
		if friendCardIndex > maxCardIndex {
			maxCardIndex = friendCardIndex
			friendIndex = index
		}
	}
	return
}

func FindMyCardMinIndex(friendIndex int, friendsCards map[int]int, myCards map[int]struct{}) (int, error) {
	minIndex := 0x7fffffffffffffff
	for index := range myCards {
		if index < minIndex && index > friendsCards[friendIndex] {
			minIndex = index
		}
	}
	if minIndex == 0x7fffffffffffffff {
		return 0, fmt.Errorf("")
	}
	return minIndex, nil
}

func NewSolution(in io.Reader, out io.Writer) {
	var friendsCount, myCardsCount int
	fmt.Fscan(in, &friendsCount, &myCardsCount)

	friendsCards := make(map[int]int)
	for friendIndex := 0; friendIndex < friendsCount; friendIndex++ {
		var cardsAmount int
		fmt.Fscan(in, &cardsAmount)
		friendsCards[friendIndex] = cardsAmount
	}

	FriendsGifts := make([]int, 0, friendsCount)

	err := FindCards(friendsCards, &FriendsGifts, myCardsCount, friendsCount)
	if err != nil {
		fmt.Println(-1)
		return
	}

	for _, value := range FriendsGifts {
		fmt.Fprint(out, value)
		fmt.Fprint(out, " ")
	}
	fmt.Fprintln(out)

}

func FindCards(friendsCards map[int]int, FriendsGifts *[]int, myCardsCount int, frinedsCount int) error {
	for friendIndex := 0; friendIndex < frinedsCount; friendIndex++ {
		proposedCard := friendsCards[friendIndex] + 1
		for {
			if proposedCard > myCardsCount {
				return fmt.Errorf("")
			}
			if CheckCardAvailable(proposedCard, FriendsGifts) {
				(*FriendsGifts) = append((*FriendsGifts), proposedCard)
				break
			}
			proposedCard++
		}
	}
	return nil
}

func CheckCardAvailable(card int, FriendsGifts *[]int) bool {
	for _, value := range *FriendsGifts {
		if value == card {
			return false
		}
	}
	return true
}

type FriendGift struct {
	pos  int
	card int
}

func TheNewesSolution(in io.Reader, out io.Writer) {
	var friendsCount, myCardsCount int
	fmt.Fscan(in, &friendsCount, &myCardsCount)

	friendsCards := make(map[int]int)
	FriendsGifts := make([]FriendGift, 0, friendsCount)
	for friendIndex := 1; friendIndex <= friendsCount; friendIndex++ {
		var cardsAmount int
		fmt.Fscan(in, &cardsAmount)
		if cardsAmount == myCardsCount {
			fmt.Fprintln(out, -1)
			return
		}
		friendsCards[friendIndex] = cardsAmount
		FriendsGifts = append(FriendsGifts, FriendGift{
			pos:  friendIndex,
			card: cardsAmount + 1})
	}

	err := ReshuffleCards(&FriendsGifts, myCardsCount)
	if err != nil {
		fmt.Fprintln(out, -1)
		return
	}

	for _, friendGift := range FriendsGifts {
		friendsCards[friendGift.pos] = friendGift.card
	}

	for friendIndex := 1; friendIndex <= friendsCount; friendIndex++ {
		fmt.Fprint(out, friendsCards[friendIndex])
		fmt.Fprint(out, " ")
	}
	fmt.Fprintln(out, "")

}

func ReshuffleCards(FriendsGifts *[]FriendGift, myCardsCount int) error {
	sort.Slice(*FriendsGifts, func(i, j int) bool {
		return (*FriendsGifts)[i].card < (*FriendsGifts)[j].card
	})

	for index, friendGift := range *FriendsGifts {
		if index == 0 {
			continue
		}
		if friendGift.card == (*FriendsGifts)[index-1].card || friendGift.card < (*FriendsGifts)[index-1].card {
			(*FriendsGifts)[index].card = (*FriendsGifts)[index-1].card + 1

		}
		if (*FriendsGifts)[index].card > myCardsCount {
			return fmt.Errorf("")
		}
	}
	return nil
}
