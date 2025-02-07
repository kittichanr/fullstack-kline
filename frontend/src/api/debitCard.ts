import api from "./client"

export const fetchDebitCardData = async (userId: string) => {
  const response = await api.get(
    `${import.meta.env.VITE_API_URL}/debit-card/${userId}`
  )

  return response.data
}
