import (
    "fmt"
    "github.com/WireGuard/wireguard-go/wgctrl"
)

func createWgInterface() {
    interface, err := wgctrl.NewInterface("my-interface")
    if err != nil {
        fmt.Println(err)
        return
    }
}

func setInterfacePrivateKey(interface wgctrl.Interface, string privateKey) {
    err = interface.SetPrivateKey([]byte(privateKey))
    if err != nil {
        fmt.Println(err)
        return
    }
}

func addInterfaceToNetwork(interface wgctrl.Interface) {
    err = interface.AddToNetwork("my-network")
    if err != nil {
        fmt.Println(err)
        return
    }
}

func startInterface(interface wgctrl.Interface) {
    err = interface.Start()
    if err != nil {
        fmt.Println(err)
        return
    }
}

func WaitForUp() {
    err = interface.WaitForUp()
    if err != nil {
        fmt.Println(err)
        return
    }

    // Your implant is now running.
    fmt.Println("Your WireGuard implant is now running.")
}