import axios from "axios"

const api = axios.create({
  baseURL: "http://localhost:8080",
})

api.interceptors.request.use((config) => {
  if (config.url?.includes("/api/auth-pin")) {
    return config
  }

  const token = sessionStorage.getItem("token")
  if (!token) {
    throw new Error("User not authenticated")
  } else {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      sessionStorage.removeItem("token")
      window.location.href = "/pin"
    }
    return Promise.reject(error)
  }
)

export default api
