import Head from "next/head";
import { useQuery } from "@apollo/client";
import { GetUsersDocument, GetUsersQuery } from "../graph/dist/client";
import styles from "../styles/Home.module.css";

export default function Home() {
  const { data } = useQuery<GetUsersQuery>(GetUsersDocument);

  return (
    <div className={styles.container}>
      <Head>
        <title>Nulledge</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main className={styles.main}>
        <h1 className={styles.title}>Nulledge Network</h1>

        <p className={styles.description}>nulledge is notebook system.</p>

        <div className={styles.grid}>
          <div className={styles.card}>
            <h3>Users</h3>
            {data?.users?.map((u) => (
              <div key={u.id}>
                <p>{u.name}</p>
              </div>
            ))}
          </div>
          <div className={styles.card}>
            <h3>Notes</h3>
          </div>
          <div className={styles.card}>
            <h3>Configuration</h3>
          </div>
        </div>
      </main>
      <footer className={styles.footer}>tomato3713, 2023</footer>
    </div>
  );
}
