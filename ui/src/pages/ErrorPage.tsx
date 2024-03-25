import { ReactNode } from "react";
import { Blankslate } from '@primer/react/drafts'
import { ArrowLeftIcon, CircleSlashIcon } from "@primer/octicons-react";


type Props = { children: ReactNode; message?: string }

export const ErrorPage = ({ children }: Props) => {
    return (
        <Blankslate>
            <Blankslate.Visual>
                <CircleSlashIcon size={"large"} />
            </Blankslate.Visual>
            {children}
            <Blankslate.PrimaryAction href="/">
                <ArrowLeftIcon />
                Back Home
            </Blankslate.PrimaryAction>
        </Blankslate>
    )
}

export const NotFound = () => {
    return (
        <ErrorPage>
            <Blankslate.Heading>   Page Not Found</Blankslate.Heading>
            <Blankslate.Description>
                The page you are looking for doesn&apos;t exist or you may not have access to it.
            </Blankslate.Description>
        </ErrorPage>
    )
}
