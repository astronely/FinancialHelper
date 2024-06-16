import {createContext, useState} from "react";
import axios from "axios";
import {toast} from "react-toastify";

export const AppContext = createContext({
    wallets: [],
    setWallets: null,
    updateWallets: null,
    getWallet: null,
})

export const AppProvider = ({ children }) => {
    const [wallets, setWallets] = useState([])

    const getWallet = wallet => {
        // console.log("getWallet:", wallet)
        setWallets(wallets => [...wallets, {
            id: wallet.ID,
            name: wallet.Name,
            value: wallet.Balance.Float64,
            currency: wallet.Currency
        }])
    }

    const updateWallets = async () => {
        // if (wallets.length === 0) {
        await axios.get("http://localhost:8080/api/wallets/get_wallets", {withCredentials: true})
            .then(response => {
                setWallets([])
                // console.log("UPDATE WALLETS")
                if (response.data === null) return
                if (response.data.sort((a, b) => a.ID - b.ID)) {
                    for (let wallet of response.data) {
                        getWallet(wallet)
                        // console.log(wallet)
                    }
                }
            })
            .catch(error => {
                console.log(error)
                toast.error("Не получилось загрузить кошельки")
            })
        // }
    }

    return (
        <AppContext.Provider value={{
            wallets, setWallets, updateWallets, getWallet
        }}>
            {children}
        </AppContext.Provider>
    )
}