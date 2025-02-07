import api from "./client"

export const fetchBannerData = async (userId: string) => {
  const response = await api.get(
    `${import.meta.env.VITE_API_URL}/banner/${userId}`
  )

  return response.data
}
