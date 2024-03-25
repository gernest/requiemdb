import { createRoutesFromElements, Route } from "react-router-dom";
import { NotFound, RouterDataErrorBoundary, Layout } from "./pages";

export const routes = createRoutesFromElements(
    <Route element={(<Layout />)} errorElement={<RouterDataErrorBoundary />}>
        <Route path="*" element={<NotFound />} />
    </Route>
)