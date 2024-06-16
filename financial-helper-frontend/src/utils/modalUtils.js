export function openModal(setIsActive, setModal, modalName) {
    // console.log("OPEN ADD EXPENSE")
    setIsActive(true)
    setModal(modalName)
}

export function openConfirm(setIsActive, setModal, modalName, item, setItemToDelete) {
    openModal(setIsActive, setModal, modalName)
    setItemToDelete(item)
}

export function isPriceCorrect(price) {
    return /^\d+(\.\d{1,2})?$/.test(price)
}
