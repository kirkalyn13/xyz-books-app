import Header from './components/Header/Header'
import Footer from './components/Footer/Footer'
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom"
import BookList from './components/BookList/BookList'
import Book from './components/Book/Book';

const router = createBrowserRouter([
  {
      path: "/",
      element: <BookList />,
  },
  {
      path: "/books/:isbn13",
      element: <Book />,
  },
]);

function App() {
  return (
    <main>
      <Header />
        <RouterProvider router={router} />
      <Footer />
    </main>
  )
}

export default App
