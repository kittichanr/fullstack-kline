import api from "./client"

export const fetchAccountData = async (userId: string) => {
  const response = await api.get(
    `${import.meta.env.VITE_API_URL}/account/${userId}`
  )

  return response.data?.accounts
}
