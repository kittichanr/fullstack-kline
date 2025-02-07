export const thBaht = new Intl.NumberFormat("th-TH", {
  style: "currency",
  currency: "THB",
})

export const maskCardNumber = (cardNumber: string) => {
  return cardNumber.replace(/.(?=.{4})/g, "•")
}
