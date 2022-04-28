import React from 'react';
import clsx from 'clsx';
import Layout from '@theme/Layout';
import Link from '@docusaurus/Link';
import styles from './index.module.css';
import HomepageFeatures from '@site/src/components/HomepageFeatures';

function HomepageHeader() {
	return (
		<header className={clsx('hero hero--primary', styles.heroBanner)}>
			<div className="container">
				<h1 className="hero__title">Kitabisa</h1>
				<p className="hero__subtitle">test documentation</p>
				<div className={styles.buttons}>
					<Link
						className="button button--secondary button--lg"
						to="/docs/problem-3/question">
						Go To Documentation 📄
					</Link>
				</div>
			</div>
		</header>
	);
}

export default function Home() {
	return (
		<Layout>
			<HomepageHeader/>
			<main>
				<HomepageFeatures/>
			</main>
		</Layout>
	);
}
