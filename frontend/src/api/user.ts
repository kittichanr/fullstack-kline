import api from "./client"

export const fetchUserData = async (userId: string) => {
  const response = await api.get(
    `${import.meta.env.VITE_API_URL}/user/greeting/${userId}`
  )

  return response.data
}
