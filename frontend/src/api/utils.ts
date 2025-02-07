import { jwtDecode } from "jwt-decode"

export const isTokenExpired = () => {
  const token = sessionStorage.getItem("token")
  if (!token) return true

  const decoded = jwtDecode(token) as any
  return decoded.exp * 1000 < Date.now()
}

export const getUserId = () => {
  const token = sessionStorage.getItem("token")
  if (!token) return ""

  const decoded = jwtDecode(token) as any
  return decoded?.user_id ?? ""
}
