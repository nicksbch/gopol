document.addEventListener('DOMContentLoaded', function() {
    const width = document.getElementById('graph').clientWidth;
    const height = window.innerHeight;

    const svg = d3.select('#graph')
        .append('svg')
        .attr('width', width)
        .attr('height', height);

    // Create the force simulation
    const simulation = d3.forceSimulation(graphData.People)
        .force('link', d3.forceLink(graphData.Relationships)
            .id(d => d.id)
            .distance(150))
        .force('charge', d3.forceManyBody().strength(-300))
        .force('center', d3.forceCenter(width / 2, height / 2));

    // Draw the relationships
    const links = svg.append('g')
        .selectAll('line')
        .data(graphData.Relationships)
        .enter()
        .append('line')
        .attr('class', 'link')
        .attr('stroke-width', d => getStrokeWidth(d.strength));

    // Draw the nodes
    const nodes = svg.append('g')
        .selectAll('g')
        .data(graphData.People)
        .enter()
        .append('g')
        .attr('class', 'node')
        .call(d3.drag()
            .on('start', dragstarted)
            .on('drag', dragged)
            .on('end', dragended));

    nodes.append('circle')
        .attr('r', 20)
        .attr('fill', d => getNodeColor(d.level));

    nodes.append('text')
        .text(d => d.name)
        .attr('x', 25)
        .attr('y', 5);

    // Update positions on each tick
    simulation.on('tick', () => {
        links
            .attr('x1', d => d.source.x)
            .attr('y1', d => d.source.y)
            .attr('x2', d => d.target.x)
            .attr('y2', d => d.target.y);

        nodes
            .attr('transform', d => `translate(${d.x},${d.y})`);
    });

    // Handle node clicks
    nodes.on('click', (event, d) => {
        showDetails(d);
    });

    // Helper functions
    function getNodeColor(level) {
        const colors = {
            'contractor': '#ff7f0e',
            'employee': '#1f77b4',
            'senior employee': '#2ca02c',
            'manager': '#d62728',
            'executive': '#9467bd'
        };
        return colors[level] || '#999';
    }

    function getStrokeWidth(strength) {
        const widths = {
            'weak': 1,
            'medium': 2,
            'strong': 3
        };
        return widths[strength] || 1;
    }

    function showDetails(node) {
        const details = document.getElementById('node-details');
        details.innerHTML = `
            <h3>${node.name}</h3>
            <p>Role: ${node.role}</p>
            <p>Team: ${node.team}</p>
            <p>Level: ${node.level}</p>
            <p>Influence: ${node.influence}</p>
        `;
    }

    // Drag functions
    function dragstarted(event) {
        if (!event.active) simulation.alphaTarget(0.3).restart();
        event.subject.fx = event.subject.x;
        event.subject.fy = event.subject.y;
    }

    function dragged(event) {
        event.subject.fx = event.x;
        event.subject.fy = event.y;
    }

    function dragended(event) {
        if (!event.active) simulation.alphaTarget(0);
        event.subject.fx = null;
        event.subject.fy = null;
    }
});