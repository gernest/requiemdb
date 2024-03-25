import { ErrorBoundary as BaseErrorBoundary } from 'react-error-boundary'
import { useRouteError } from 'react-router-dom'
import { RpcError } from "@protobuf-ts/runtime-rpc";
import { ErrorPage, NotFound } from "./ErrorPage";
import { ReactNode } from 'react';
import { Blankslate } from '@primer/react/drafts'

type Props = { error: Error | RpcError }

const ErrorFallback = ({ error }: Props) => {
    if ('status' in error && error.status === 404) {
        return <NotFound />
    }
    if ('code' in error) {
        const e = error as RpcError;
        return (
            <ErrorPage>
                <Blankslate.Heading>
                    {e.code}
                </Blankslate.Heading>
                <Blankslate.Description>
                    {e.toString()}
                </Blankslate.Description>
            </ErrorPage>
        )
    }

    return (
        <ErrorPage>
            <Blankslate.Heading>
                Something went wrong
            </Blankslate.Heading>
            <Blankslate.Description>
                {error.message}
            </Blankslate.Description>
        </ErrorPage>
    )
}

export const ErrorBoundary = (props: { children: ReactNode }) => (
    <BaseErrorBoundary FallbackComponent={ErrorFallback}>{props.children}</BaseErrorBoundary>
)

export function RouterDataErrorBoundary() {
    const error = useRouteError() as Props['error']
    return <ErrorFallback error={error} />
}