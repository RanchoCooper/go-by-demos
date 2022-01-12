package main

/**
 * @author Rancho
 * @date 2022/1/11
 */

func main() {
    welcomer := DefaultContainer.GetCustomerWelcome()

    _ = welcomer.Welcome("cooper", "cooper@gmail.com")
}
