import { Route, Routes } from "react-router"
import Splash from "./pages/Splash"
import Pin from "./pages/Pin"
import MainBank from "./pages/MainBank"
import { QueryClient, QueryClientProvider } from "react-query"

function App() {
  const queryClient = new QueryClient()

  return (
    <QueryClientProvider client={queryClient}>
      <Routes>
        <Route path="/" element={<Splash />} />
        <Route path="/pin" element={<Pin />} />
        <Route path="/main-bank" element={<MainBank />} />
      </Routes>
    </QueryClientProvider>
  )
}

export default App
